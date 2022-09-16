import {defineConfig} from 'vite'
import vue from '@vitejs/plugin-vue'
import purgecss from 'rollup-plugin-purgecss';

// https://vitejs.dev/config/
export default defineConfig({
    plugins: [
        vue(),
        purgecss({
            content: [`./src/**/*.vue`],
            variables: false,
            safelist: {
                standard: [
                    /(autv|lnil|lnir|fas?)/,
                    /-(leave|enter|appear)(|-(to|from|active))$/,
                    /^(?!(|.*?:)cursor-move).+-move$/,
                    /^router-link(|-exact)-active$/,
                    /data-v-.*/,
                ],
            },
            defaultExtractor(content) {
                const contentWithoutStyleBlocks = content.replace(
                    /<style[^]+?<\/style>/gi,
                    ''
                );
                return (
                    contentWithoutStyleBlocks.match(/[A-Za-z0-9-_/:]*[A-Za-z0-9-_/]+/g) ||
                    []
                );
            },
        }),
    ],
    resolve: {
        alias: [
            {
                find: '/@src/',
                replacement: `/src/`,
            }
        ],
    },
    build: {
        chunkSizeWarningLimit: 2000,
        minify: false,
        rollupOptions: {
            output: {
                entryFileNames: `assets/[name].js`,
                chunkFileNames: `assets/[name].js`,
                assetFileNames: `assets/[name].[ext]`
            }
        }
    }
})
