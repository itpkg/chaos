require('./main.css')

import i18next from 'i18next'
// import XHR from 'i18next-xhr-backend'
import LanguageDetector from 'i18next-browser-languagedetector'

import main from './main'
import {LOCALE} from './constants'
import root from './engines'

i18next
    // .use(XHR)
    .use(LanguageDetector)
    .init({
      resources: root.locales(),
      // backend: {
      //   loadPath: process.env.CHAOS.backend + '/locales/{{lng}}',
      //   crossDomain: true
      // },
      detection: {
        order: ['querystring', 'localStorage', 'cookie', 'navigator'],
        lookupQuerystring: LOCALE,
        lookupCookie: LOCALE,
        lookupLocalStorage: LOCALE,

        caches: ['localStorage', 'cookie'],
        cookieMinutes: 365 * 24 * 60
      }
    },
    (e, t) => {
      console.log('lang: ' + i18next.language)
      main('root')
    }
      )
