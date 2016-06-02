package com.itpkg.core.oauth2;

import com.google.api.client.auth.oauth2.Credential;
import com.google.api.client.googleapis.auth.oauth2.GoogleAuthorizationCodeFlow;
import com.google.api.client.googleapis.auth.oauth2.GoogleTokenResponse;
import com.google.api.client.http.*;
import com.google.api.client.http.javanet.NetHttpTransport;
import com.google.api.client.json.JsonFactory;
import com.google.api.client.json.JsonObjectParser;
import com.google.api.client.json.jackson2.JacksonFactory;
import com.google.api.services.oauth2.model.Userinfoplus;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Component;

import javax.annotation.PostConstruct;
import java.io.IOException;
import java.util.ArrayList;
import java.util.List;

/**
 * Created by flamen on 16-6-2.
 */
@Component
public class Oauth2GoogleHelper {
    public Userinfoplus getUser(String code) throws IOException {
        HttpResponse response = request(
                "https://www.googleapis.com/oauth2/v3/userinfo",
                code).execute();
        try {
            Userinfoplus u = response.parseAs(Userinfoplus.class);
            logger.debug("GOOGLE USER {}", u.toPrettyString());
            return u;
        } finally {
            response.disconnect();
        }
    }

    public String url() {
        return flow.newAuthorizationUrl()
                .setRedirectUri(redirectUrl)
                .setState(ID + state).build();
    }

    private HttpRequest request(String url, String code) throws IOException {
        GoogleTokenResponse response = flow.newTokenRequest(code).setRedirectUri(redirectUrl).execute();
        Credential credential = flow.createAndStoreCredential(response, null);
        JsonObjectParser parser = new JsonObjectParser(factory); //flow.getJsonFactory().createJsonObjectParser();

        HttpRequestFactory factory = transport.createRequestFactory(credential);
        HttpRequest request = factory.buildGetRequest(new GenericUrl(url));
        request.getHeaders().setContentType("application/json");
        request.setParser(parser);

        return request;
    }

    @PostConstruct
    void init() {
        scopes = new ArrayList<>();
        scopes.add("https://www.googleapis.com/auth/userinfo.profile");
        scopes.add("https://www.googleapis.com/auth/userinfo.email");

        transport = new NetHttpTransport();
        factory = new JacksonFactory();

        flow = new GoogleAuthorizationCodeFlow.Builder(
                transport,
                factory,
                clientId,
                clientSecret,
                scopes).build();
    }


    GoogleAuthorizationCodeFlow flow;
    List<String> scopes;
    HttpTransport transport;
    JsonFactory factory;

    @Value("${oauth.google.redirectUrl}")
    String redirectUrl;
    @Value("${oauth.google.clientId}")
    String clientId;
    @Value("${oauth.google.clientSecret}")
    String clientSecret;
    @Value("${oauth.google.state}")
    String state;


    public final static String ID = "google-";
    private final static Logger logger = LoggerFactory.getLogger(Oauth2GoogleHelper.class);

}
