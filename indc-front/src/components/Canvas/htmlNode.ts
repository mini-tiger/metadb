import { HtmlNodeModel, HtmlNode } from '@logicflow/core'
class ZHModel extends HtmlNodeModel {
  setAttributes() {
    this.text.editable = false // 禁止节点文本编辑
    // 设置节点宽高和锚点
    const width = 36
    const height = 36
    this.width = width
    this.height = height
    this.anchorsOffset = [
      [width / 2, 0],
      [0, height / 2],
      [-width / 2, 0],
      [0, -height / 2],
      // [0, 0],
    ]
  }
  getTextStyle() {
    const style = super.getTextStyle()
    style.fontSize = this.properties.fontSize || 12
    style.color = '#302E6D'
    style.fontWeight = 500
    style.fontFamily = 'PingFangSC-Medium, PingFang SC'
    return style
  }
}
class ZHNode extends HtmlNode {
  setHtml(rootEl: HTMLElement) {
    const { properties } = this.props.model
    const el = document.createElement('div')
    el.className = 'view-wrapper'
    const html = `
       <div class="node-content" id=${properties.id}></div><div class="content-text">${properties.text}</div>
    `
    el.innerHTML = html
    // 需要先把之前渲染的子节点清除掉。
    rootEl.innerHTML = ''
    rootEl.appendChild(el)
  }
}
export default {
  type: 'htmlNode',
  view: ZHNode,
  model: ZHModel,
}
