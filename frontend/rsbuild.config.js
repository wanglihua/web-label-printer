import {defineConfig} from '@rsbuild/core';
import {pluginVue} from '@rsbuild/plugin-vue';
import {pluginBabel} from '@rsbuild/plugin-babel';
import {pluginImageCompress} from '@rsbuild/plugin-image-compress';
import {pluginSass} from '@rsbuild/plugin-sass';
import {pluginVueJsx} from '@rsbuild/plugin-vue-jsx';

// Docs: https://rsbuild.rs/config/
export default defineConfig(({env, command, envMode}) => {
    return {
        dev: {
            hmr: true,
        },
        html: {
            template: "./index.html",
        },
        output: {
            target: 'web',
            cleanDistPath: env === 'production',
            filenameHash: 'contenthash:8',
            distPath: {
                // root: './dist',
                root: '../backend/route/static_route/www_root',
                html: './',
                js: 'static/js',
                css: 'static/css',
                font: 'static/font',
                image: 'static/image',
                assets: 'static/assets',
            },
            polyfill: 'off',
            minify: true,
            copy: [
                {from: "./public/favicon.webp", to: "favicon.webp"},
            ],
        },
        performance: {
            chunkSplit: {
                strategy: 'split-by-experience',
            },
            removeConsole: true,
            // bundleAnalyze: {},
        },
        plugins: [
            pluginVue(),
            pluginBabel({
                include: /\.(?:jsx|tsx)$/,
            }),
            pluginVueJsx(),
            pluginSass(),
            pluginImageCompress(),
        ],
        server: {
            base: '/',
            open: true,
            port: 3000,
            printUrls: true,
            proxy: {
                '/api': 'http://localhost:8113',
                '/print': 'http://localhost:8113',
                '/websocket': {
                    target: 'http://localhost:8113',
                    ws: true,
                },
            },
            publicDir: {
                copyOnBuild: false,
                name: './public',
            },
        },
        source: {
            entry: {
                index: './src/index.js',
            },

        },
        resolve: {
            alias: {
                '@': './src',
            },
        },
    };
});
