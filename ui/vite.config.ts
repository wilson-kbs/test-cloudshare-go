import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import copy from 'rollup-plugin-copy'

import * as path from 'path'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    vue(),
    copy({
      targets: [
        { src: 'public/**/*', dest: 'dist' },
        { src: 'src/assets/**/*', dest: 'dist/assets' }
      ],
      hook: 'buildStart',
      verbose: true
    })
  ],
  root: "src",
  build: {
    outDir: path.resolve(__dirname, 'dist'),
    emptyOutDir: true,
    // generate manifest.json in outDir
    manifest: true,
    rollupOptions: {
      input: '/main.ts',
    }
  },
  server: {
    // required to load scripts from custom host
    cors: true,

    // we need a strict port to match on PHP side
    // change freely, but update on PHP to match the same port
    strictPort: true,
    port: 3000
  },
  // resolve: {
  //   alias: {
  //     vue: 'vue/dist/vue.esm-bundler.js'
  //   }
  // }
})
