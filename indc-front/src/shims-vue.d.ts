/* eslint-disable */
declare module '*.vue' {
  import type { DefineComponent } from 'vue'
  const component: DefineComponent<{}, {}, any>

  export default component
}
declare module '@/components/svgIcon' {}
declare module '@/directives/directive/watermark' {}
declare module '@/store/formula.json' {}
declare module '@/assets/style/var.scss' {
  export const mainColor: string
  export const btnTextColor: string
  export const baseColor: string
}
