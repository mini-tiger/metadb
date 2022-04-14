import Breadcrum from './Breadcrumb/index.vue'
import SvgIcon from '../assets/icons'
import ZHButton from '@/components/Button/Button.vue'
import ZHSelect from '@/components/Select/index.vue'
import ZHDatePicker from '@/components/Date/DatePicker/index.vue'
import ZHDropdownTree from '@/components/DropdownTree/index.vue'
import ZHSelectAndInput from '@/components/SelectAndInput/index.vue'
import ZHTable from '@/components/Table/Table.vue'
import ZHSwitch from '@/components/Switch/index.vue'

const components = {
  Breadcrum,
  SvgIcon,
  ZHButton,
  ZHSelect,
  ZHDatePicker,
  ZHDropdownTree,
  ZHSelectAndInput,
  ZHTable,
  ZHSwitch,
}
export default {
  install(app) {
    Object.keys(components).forEach((key) => {
      app.component(key, components[key])
    })
  },
}
