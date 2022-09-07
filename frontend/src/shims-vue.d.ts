declare module '*.vue' {
    import { defineComponent } from 'vue'
    const Component: ReturnType<typeof defineComponent>
    export default Component
}


declare module 'nprogress'
declare module 'simplebar-vue'
declare module '@vueform/multiselect'
declare module '@vueform/slider'
declare module 'vue-accessible-color-picker'