import Vue from 'vue'

// import './styles/quasar.sass'
import iconSet from 'quasar/icon-set/fontawesome-v5'
import lang from 'quasar/lang/zh-hans.js'
import '@quasar/extras/ionicons-v4/ionicons-v4.css'
import { Quasar } from 'quasar'

Vue.use(Quasar, {
  config: {
    animations: 'all'
  },
  plugins: {
  },
  lang: lang,
  iconSet: iconSet
 })