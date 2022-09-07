import { defineConfig } from 'vite';
import vue from '@vitejs/plugin-vue';
import purgecss from 'rollup-plugin-purgecss';
import Pages from 'vite-plugin-pages';

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
      vue({
        include: [/\.vue$/],
      }),

    /**
     * vite-plugin-pages plugin generate routes based on file system
     *
     * @see https://github.com/hannoeru/vite-plugin-pages
     */
    Pages({
      pagesDir: [
        {
          dir: 'src/pages',
          baseRoute: '',
        },
      ],
    }),

    /**
     * rollup-plugin-purgecss plugin is responsible of purging css rules
     * that are not used in the bundle
     *
     * @see https://github.com/FullHuman/purgecss/tree/main/packages/rollup-plugin-purgecss
     */
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
