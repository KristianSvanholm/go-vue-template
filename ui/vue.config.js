const { defineConfig } = require('@vue/cli-service')
module.exports = defineConfig({
    transpileDependencies: true,
    assetsDir: "static",
    publicPath: "/",
    pluginOptions: {
        vuetify: {
			// https://github.com/vuetifyjs/vuetify-loader/tree/next/packages/vuetify-loader
		}
    },
    devServer: {
        proxy: {
            "^/api": {
                target: 'https://localhost:8080',
                secure: false,
                changeOrigin: true,
            },
        },
    },
})
