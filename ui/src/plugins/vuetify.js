// Styles
import '@mdi/font/css/materialdesignicons.css'
import 'vuetify/styles'

// Vuetify
import { createVuetify } from 'vuetify'
//import { VDataTable } from 'vuetify/labs/VDataTable'

let theme = localStorage.getItem('theme')
if (theme === null) {
    theme = "light"
    localStorage.setItem('theme', 'light')
}

export default createVuetify({
    theme:{
        defaultTheme: theme,
    },
    components: {
        //VDataTable
    }
})
