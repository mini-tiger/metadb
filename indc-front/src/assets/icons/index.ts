import SvgIcon from '@/components/svgIcon/index.vue'

const requireAll = (requireContext) => requireContext.keys().map(requireContext)
const req = require.context('./svg', true, /\.svg$/)
requireAll(req)
export default SvgIcon
