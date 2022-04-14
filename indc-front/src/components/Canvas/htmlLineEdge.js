import { LineEdge, PolylineEdgeModel } from '@logicflow/core'
import style from '@/assets/style/var.scss'
const activeColor = style.mainColor
const defaultColor = style.baseColor
class LineModel extends PolylineEdgeModel {
  customTextPosition = true
  getTextPosition() {
    const position = super.getTextPosition()
    const pointsList = []

    this.pointsList &&
      this.pointsList.forEach((item) => {
        pointsList.push({ x: Number(item.x), y: Number(item.y) })
      })

    position.x =
      pointsList[0].x +
      (pointsList[pointsList.length - 1].x - pointsList[0].x) / 2
    position.y =
      pointsList[0].y +
      (pointsList[pointsList.length - 1].y - pointsList[0].y) / 2

    return position
  }
  getEdgeStyle() {
    const style = super.getEdgeStyle()
    const { properties } = this
    // style.strokeWidth = 1
    if (properties.isActived) {
      style.stroke = activeColor
    } else {
      style.stroke = defaultColor
    }
    return style
  }
  getTextStyle() {
    const style = super.getTextStyle()
    const { properties } = this
    if (properties.isActived) {
      style.color = activeColor
    } else {
      style.color = defaultColor
    }
    style.fontSize = 12
    style.fontWeight = 500
    return style
  }
  getOutlineStyle() {
    const style = super.getOutlineStyle()
    style.stroke = defaultColor
    style.hover.stroke = defaultColor
    return style
  }
}

export default {
  type: 'htmlLineEdge',
  view: LineEdge,
  model: LineModel,
}
