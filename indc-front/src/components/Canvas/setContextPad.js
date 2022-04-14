import { ElMessage } from 'element-plus'
import store from '@/store'
import * as types from '@/store/action-types'
const setContextPad = (lf) => {
  const deleteConfig = {
    icon: require('../../assets/icons/png/canvas-delete.png'),
    callback: (data) => {
      let incomingEdge = lf.getNodeIncomingEdge(data.id).length / 2
      let outgoingEdge = lf.getNodeOutgoingEdge(data.id).length / 2
      if (incomingEdge > 0 || outgoingEdge > 0) {
        return ElMessage.error('存在相关的边，无法删除')
      }
      lf.deleteElement(data.id)
      store.dispatch(types.DELETE_POINT, data.id)
      lf.extension.contextPad.hideContextMenu()
    },
  }
  lf.extension.contextPad.setContextMenuItems([deleteConfig])
}

export default setContextPad
