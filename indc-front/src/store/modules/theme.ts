import { getItem, setItem } from '@/utils/commUtils'
import { MAIN_COLOR, DEFAULT_COLOR, SET_MAIN_COLOR } from '@/store/action-types'
import variables from '@/assets/style/var.scss'
export default {
  state: () => ({
    themeColor: getItem(MAIN_COLOR) || DEFAULT_COLOR,
    ...variables,
  }),
  mutations: {
    /**
     * 设置主题色
     */
    [SET_MAIN_COLOR](state, newColor) {
      state.themeColor = newColor
      setItem(MAIN_COLOR, newColor)
    },
  },
}
