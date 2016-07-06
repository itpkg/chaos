package mail

import (
	"fmt"
	"os"
	"path"
	"text/template"

	"github.com/itpkg/chaos/engines/ops"
	"github.com/itpkg/chaos/web"
	"github.com/spf13/viper"
	"github.com/urfave/cli"
)

func write_config(c *cli.Context, files map[string]string) error {
	user := c.String("user")
	password := c.String("password")
	if len(user) == 0 || len(password) == 0 {
		cli.ShowSubcommandHelp(c)
		return nil
	}
	for n, b := range files {
		if err := os.MkdirAll(path.Dir(n), 0700); err != nil {
			return err
		}
		fmt.Printf("generate file %s\n", n)
		fd, err := os.OpenFile(
			n,
			os.O_WRONLY|os.O_CREATE|os.O_EXCL,
			0600)
		if err != nil {
			return err
		}
		defer fd.Close()
		t, err := template.New("").Parse(b)
		if err != nil {
			return err
		}
		driver := viper.GetString("database.driver")
		switch driver {
		case "postgres":
			driver = "pgsql"
		}
		args := viper.GetStringMapString("database.args")
		if err = t.Execute(fd, struct {
			Driver      string
			User        string
			Password    string
			Name        string
			Host        string
			Port        string
			UserTable   string
			AliasTable  string
			DomainTable string
		}{
			Driver:      driver,
			Host:        args["host"],
			Name:        args["dbname"],
			Port:        args["port"],
			Password:    password,
			User:        user,
			UserTable:   User{}.TableName(),
			AliasTable:  Alias{}.TableName(),
			DomainTable: Domain{}.TableName(),
		}); err != nil {
			return err
		}
	}

	return nil
}

//Shell command
func (p *Engine) Shell() []cli.Command {
	return []cli.Command{
		{
			Name:    "ops:mail",
			Aliases: []string{"ops:m"},
			Usage:   "mail server manage",
			Subcommands: []cli.Command{
				{
					Name:    "postfix",
					Aliases: []string{"pf"},
					Flags:   []cli.Flag{ops.FlagDatabaseUser, ops.FlagDatabasePassword},
					Usage:   "generate postfix config files",
					Action: web.Action(func(c *cli.Context) error {
						files := map[string]string{
							"etc/postfix/main.cf": `
# See /usr/share/postfix/main.cf.dist for a commented, more complete version

# Debian specific:  Specifying a file name will cause the first
# line of that file to be used as the name.  The Debian default
# is /etc/mailname.
#myorigin = /etc/mailname

smtpd_banner = $myhostname ESMTP $mail_name (IT-PACKAGE)
biff = no

# appending .domain is the MUA's job.
append_dot_mydomain = no

# Uncomment the next line to generate "delayed mail" warnings
#delay_warning_time = 4h

readme_directory = no

smtpd_tls_cert_file=/etc/dovecot/dovecot.pem
smtpd_tls_key_file=/etc/dovecot/private/dovecot.pem
smtpd_use_tls=yes
smtpd_tls_auth_only = yes

#Enabling SMTP for authenticated users, and handing off authentication to Dovecot
smtpd_sasl_type = dovecot
smtpd_sasl_path = private/auth
smtpd_sasl_auth_enable = yes

smtpd_recipient_restrictions =
        permit_sasl_authenticated,
        permit_mynetworks,
        reject_unauth_destination

# See /usr/share/doc/postfix/TLS_README.gz in the postfix-doc package for
# information on enabling SSL in the smtp client.

myhostname = CHANGE-ME
alias_maps = hash:/etc/aliases
alias_database = hash:/etc/aliases
myorigin = /etc/mailname
mydestination = localhost
relayhost =
mynetworks = 127.0.0.0/8 [::ffff:127.0.0.0]/104 [::1]/128
mailbox_size_limit = 0
recipient_delimiter = +
inet_interfaces = all

#Handing off local delivery to Dovecot's LMTP, and telling it where to store mail
virtual_transport = lmtp:unix:private/dovecot-lmtp

#Virtual domains, users, and aliases
virtual_mailbox_domains = {{.Driver}}:/etc/postfix/virtual-mailbox-domains.cf
virtual_mailbox_maps = {{.Driver}}:/etc/postfix/virtual-mailbox-maps.cf
virtual_alias_maps = {{.Driver}}:/etc/postfix/virtual-alias-maps.cf,
  {{.Driver}}:/etc/postfix/virtual-email2email.cf
							`,
							"etc/postfix/virtual-mailbox-domains.cf": `
user = {{.User}}
password = {{.Password}}
hosts = {{.Host}}
dbname = {{.Name}}
query = SELECT 1 FROM {{.DomainTable}} WHERE name='%s'
							`,
							"etc/postfix/virtual-mailbox-maps.cf": `
user = {{.User}}
password = {{.Password}}
hosts = {{.Host}}
dbname = {{.Name}}
query = SELECT 1 FROM {{.UserTable}} WHERE email='%s'
							`,
							"etc/postfix/virtual-alias-maps.cf": `
user = {{.User}}
password = {{.Password}}
hosts = {{.Host}}
dbname = {{.Name}}
query = SELECT destination FROM {{.AliasTable}} WHERE source='%s'
							`,
							"etc/postfix/virtual-email2email.cf": `
user = {{.User}}
password = {{.Password}}
hosts = {{.Host}}
dbname = {{.Name}}
query = SELECT email FROM {{.UserTable}} WHERE email='%s'
							`,
							"etc/postfix/master.cf": `
#
# Postfix master process configuration file.  For details on the format
# of the file, see the master(5) manual page (command: "man 5 master").
#
# Do not forget to execute "postfix reload" after editing this file.
#
# ==========================================================================
# service type  private unpriv  chroot  wakeup  maxproc command + args
#               (yes)   (yes)   (yes)   (never) (100)
# ==========================================================================
smtp      inet  n       -       -       -       -       smtpd
#smtp      inet  n       -       -       -       1       postscreen
#smtpd     pass  -       -       -       -       -       smtpd
#dnsblog   unix  -       -       -       -       0       dnsblog
#tlsproxy  unix  -       -       -       -       0       tlsproxy
submission inet n       -       -       -       -       smtpd
  -o syslog_name=postfix/submission
  -o smtpd_tls_security_level=encrypt
  -o smtpd_sasl_auth_enable=yes
  -o smtpd_client_restrictions=permit_sasl_authenticated,reject
  -o milter_macro_daemon_name=ORIGINATING
smtps     inet  n       -       -       -       -       smtpd
  -o syslog_name=postfix/smtps
  -o smtpd_tls_wrappermode=yes
  -o smtpd_sasl_auth_enable=yes
  -o smtpd_client_restrictions=permit_sasl_authenticated,reject
  -o milter_macro_daemon_name=ORIGINATING
							`,
							"tmp/postfix.sh": `
#!/bin/sh
postmap -q example.com {{.Driver}}:/etc/postfix/virtual-mailbox-domains.cf
postmap -q email1@example.com {{.Driver}}:/etc/postfix/virtual-mailbox-maps.cf
postmap -q alias@example.com {{.Driver}}:/etc/postfix/virtual-alias-maps.cf

							`,
						}
						return write_config(c, files)
					}),
				},
				{
					Name:    "database",
					Aliases: []string{"db"},
					Flags:   []cli.Flag{ops.FlagDatabaseUser, ops.FlagDatabasePassword},
					Usage:   "create readonly user to database",
					Action: web.Action(func(c *cli.Context) error {
						files := map[string]string{
							"tmp/mail.sql": `
CREATE USER {{.User}} PASSWORD '{{.Password}}';
GRANT SELECT ON {{.UserTable}}, {{.DomainTable}}, {{.AliasTable}} TO {{.User}};
							`,
						}
						return write_config(c, files)
					}),
				},
				{
					Name:    "dovecot",
					Aliases: []string{"dv"},
					Flags:   []cli.Flag{ops.FlagDatabaseUser, ops.FlagDatabasePassword},
					Usage:   "generate dovecot config files",
					Action: web.Action(func(c *cli.Context) error {
						files := map[string]string{
							"etc/dovecot/dovecot.conf": `
## Dovecot configuration file

# If you're in a hurry, see http://wiki2.dovecot.org/QuickConfiguration

# "doveconf -n" command gives a clean output of the changed settings. Use it
# instead of copy&pasting files when posting to the Dovecot mailing list.

# '#' character and everything after it is treated as comments. Extra spaces
# and tabs are ignored. If you want to use either of these explicitly, put the
# value inside quotes, eg.: key = "# char and trailing whitespace  "

# Default values are shown for each setting, it's not required to uncomment
# those. These are exceptions to this though: No sections (e.g. namespace {})
# or plugin settings are added by default, they're listed only as examples.
# Paths are also just examples with the real defaults being based on configure
# options. The paths listed here are for configure --prefix=/usr
# --sysconfdir=/etc --localstatedir=/var

# Enable installed protocols
!include_try /usr/share/dovecot/protocols.d/*.protocol
protocols = imap lmtp

# A comma separated list of IPs or hosts where to listen in for connections.
# "*" listens in all IPv4 interfaces, "::" listens in all IPv6 interfaces.
# If you want to specify non-default ports or anything more complex,
# edit conf.d/master.conf.
#listen = *, ::

# Base directory where to store runtime data.
#base_dir = /var/run/dovecot/

# Name of this instance. Used to prefix all Dovecot processes in ps output.
#instance_name = dovecot

# Greeting message for clients.
#login_greeting = Dovecot ready.

# Space separated list of trusted network ranges. Connections from these
# IPs are allowed to override their IP addresses and ports (for logging and
# for authentication checks). disable_plaintext_auth is also ignored for
# these networks. Typically you'd specify the IMAP proxy servers here.
#login_trusted_networks =

# Sepace separated list of login access check sockets (e.g. tcpwrap)
#login_access_sockets =

# Show more verbose process titles (in ps). Currently shows user name and
# IP address. Useful for seeing who are actually using the IMAP processes
# (eg. shared mailboxes or if same uid is used for multiple accounts).
#verbose_proctitle = no

# Should all processes be killed when Dovecot master process shuts down.
# Setting this to "no" means that Dovecot can be upgraded without
# forcing existing client connections to close (although that could also be
# a problem if the upgrade is e.g. because of a security fix).
#shutdown_clients = yes

# If non-zero, run mail commands via this many connections to doveadm server,
# instead of running them directly in the same process.
#doveadm_worker_count = 0
# UNIX socket or host:port used for connecting to doveadm server
#doveadm_socket_path = doveadm-server

# Space separated list of environment variables that are preserved on Dovecot
# startup and passed down to all of its child processes. You can also give
# key=value pairs to always set specific settings.
#import_environment = TZ

##
## Dictionary server settings
##

# Dictionary can be used to store key=value lists. This is used by several
# plugins. The dictionary can be accessed either directly or though a
# dictionary server. The following dict block maps dictionary names to URIs
# when the server is used. These can then be referenced using URIs in format
# "proxy::<name>".

dict {
  #quota = mysql:/etc/dovecot/dovecot-dict-sql.conf.ext
  #expire = sqlite:/etc/dovecot/dovecot-dict-sql.conf.ext
}

# Most of the actual configuration gets included below. The filenames are
# first sorted by their ASCII value and parsed in that order. The 00-prefixes
# in filenames are intended to make it easier to understand the ordering.
!include conf.d/*.conf

# A config file can also tried to be included without giving an error if
# it's not found:
!include_try local.conf
							`,
							"etc/dovecot/conf.d/10-mail.conf": `
mail_location = maildir:/var/mail/vhosts/%d/%n
mail_privileged_group = mail
							`,
							"etc/dovecot/conf.d/10-auth.conf": `
disable_plaintext_auth = yes
auth_mechanisms = plain login

#!include auth-system.conf.ext
!include auth-sql.conf.ext
							`,
							"etc/dovecot/conf.d/auth-sql.conf.ext": `
passdb {
  driver = sql
  args = /etc/dovecot/dovecot-sql.conf.ext
}
userdb {
  driver = static
  args = uid=vmail gid=vmail home=/var/mail/vhosts/%d/%n
}
							`,
							"etc/dovecot/dovecot-sql.conf.ext": `
driver = {{.Driver}}
connect = host={{.Host}} dbname={{.Name}} user={{.User}} password={{.Password}}
default_pass_scheme = SHA512-CRYPT
#password_query = SELECT email as user, password FROM {{.UserTable}} WHERE email='%u';
password_query = SELECT email as user, password FROM {{.UserTable}} WHERE email=(SELECT destination FROM {{.AliasTable}} WHERE source = '%u');
							`,
							"etc/dovecot/conf.d/10-master.conf": `
service auth-worker {
  # Auth worker process is run as root by default, so that it can access
  # /etc/shadow. If this isn't necessary, the user should be changed to
  # $default_internal_user.
  user = vmail
}

service imap-login {
	inet_listener imap {
	  #port = 0
	}
	inet_listener imaps {
		port = 993
		ssl = yes
	}
}

service lmtp {
  unix_listener /var/spool/postfix/private/dovecot-lmtp {
    mode = 0600
    user = postfix
    group = postfix
   }
    # Create inet listener only if you can't use the above UNIX socket
    #inet_listener lmtp {
      # Avoid making LMTP visible for the entire internet
      #address =
      #port =
    #}
  }
}

service auth {
  # auth_socket_path points to this userdb socket by default. It's typically
  # used by dovecot-lda, doveadm, possibly imap process, etc. Its default
  # permissions make it readable only by root, but you may need to relax these
  # permissions. Users that have access to this socket are able to get a list
  # of all usernames and get results of everyone's userdb lookups.
  unix_listener /var/spool/postfix/private/auth {
    mode = 0666
    user = postfix
    group = postfix
  }

  unix_listener auth-userdb {
    mode = 0600
    user = vmail
    #group =
  }

  # Postfix smtp-auth
  #unix_listener /var/spool/postfix/private/auth {
  #  mode = 0666
  #}

  # Auth process is run as this user.
  user = dovecot
}
							`,
							"etc/dovecot/conf.d/10-ssl.conf": `
ssl_cert = </etc/dovecot/dovecot.pem
ssl_key = </etc/dovecot/private/dovecot.pem

ssl = required
							`,
							"tmp/dovecot.sh": `
#!/bin/sh
mkdir -p /var/mail/vhosts
groupadd -g 5000 vmail
useradd -g vmail -u 5000 vmail -d /var/mail
passwd -l vmail
chown -R vmail:vmail /var/mail
chown -R vmail:dovecot /etc/dovecot
chmod -R o-rwx /etc/dovecot
							`,
						}
						return write_config(c, files)
					}),
				},
			},
		},
	}
}
