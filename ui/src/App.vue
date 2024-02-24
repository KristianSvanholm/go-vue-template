<template>
    <v-app>
        <v-main>
            <v-app-bar :elevation="2" rounded>
                <v-app-bar-title >
                    <v-btn variant="plain" :to="`/`">IG Parser</v-btn>
                    <v-btn icon @click="toggleTheme"><v-icon>mdi-white-balance-sunny</v-icon></v-btn>
                </v-app-bar-title>
                <v-btn :to="'/'" >Home</v-btn>
                <v-divider inset vertical style="margin-left: 2px; margin-right: 2px;"/>
                <v-icon>mdi-person</v-icon>
            </v-app-bar>

            <br/>
            <RouterView/>
        </v-main>
    </v-app>
</template>

<style>
    .overlay {
        position: fixed;
        top: 0;
        left: 0;
        width: 100%;
        height: 100%;
        backdrop-filter: blur(1px); /* Adjust the blur amount as needed */
        z-index: 999; /* Make sure it's above other elements */
    }
</style>

<script setup> 
    import { useTheme } from 'vuetify'

    const theme = useTheme()
    function toggleTheme () {
        theme.global.name.value = theme.global.current.value.dark ? 'light' : 'dark'

        localStorage.setItem('theme', theme.global.name.value)
    }

</script>

<script>
    const debounce = (callback, delay) => {
        let tid
        return function(...args) {
            const ctx = self
            tid && clearTimeout(tid)
            tid = setTimeout(() => {
                callback.apply(ctx, args)
            }, delay)
        }
    }
    
    const _ = window.ResizeObserver
    window.ResizeObserver = class ResizeObserver extends _ {
        constructor(callback) {
            callback = debounce(callback, 20)
            super(callback)
        }
    }

    export default {
        name: 'App',

        components: {},

        data: () => ({
        
        }),

        methods: {
        
        },
        mounted: async function () {
            //Stuff that happens once 
        }
    }
</script>
