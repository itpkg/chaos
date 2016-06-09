require("bootstrap/dist/css/bootstrap.css")
require("bootstrap/dist/css/bootstrap-theme.css")
require("./main.css")

import i18next from 'i18next'
import XHR from 'i18next-xhr-backend'
import LanguageDetector from 'i18next-browser-languagedetector'

import main from './main'

i18next
    .use(XHR)
    .use(LanguageDetector)
    .init({
            backend: {
                loadPath: CHAOS_ENV.backend + '/locales/{{lng}}',
                crossDomain: true
            },
            detection: {
                order: ['querystring', 'localStorage', 'cookie', 'navigator'],
                lookupQuerystring: 'locale',
                lookupCookie: 'locale',
                lookupLocalStorage: 'locale',

                caches: ['localStorage', 'cookie'],
                cookieMinutes: 365 * 24 * 60
            }
        },
        (err, t) => {
            console.log("Lang: " + i18next.language)
            main('root')
        }
    );
