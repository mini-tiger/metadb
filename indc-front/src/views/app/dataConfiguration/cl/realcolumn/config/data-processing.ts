import { objectattr, dcCenter, dcCenterData } from './RealoumnLook'
import { arrItemSort } from '@/utils/commUtils'

export const dataProcess = () => {
  const arr = []
  let data = []
  const centerDataList = []
  for (let i = 0; i < dcCenter.length; i++) {
    const obj = {}
    obj['bk_group_id'] = dcCenter[i].bk_group_id
    obj['bk_group_index'] = dcCenter[i].bk_group_index
    obj['bk_group_name'] = dcCenter[i].bk_group_name
    for (let j = 0; j < objectattr.length; j++) {
      if (dcCenter[i].bk_group_id === objectattr[j].bk_property_group) {
        obj['bk_obj_id'] = objectattr[j].bk_obj_id
        obj['bk_property_group_name'] = objectattr[j].bk_property_group_name
        obj['bk_property_name'] = objectattr[j].bk_property_name
        obj['bk_property_type'] = objectattr[j].bk_property_type
        centerDataList.push({
          bk_property_id: objectattr[j].bk_property_id,
          bk_property_index: objectattr[j].bk_property_index,
          bk_group_id: dcCenter[i].bk_group_id,
          bk_property_name: objectattr[j].bk_property_name,
        })
        data = arrItemSort(centerDataList, 'bk_property_index', 0)
      }
    }
    const result = []
    data.map((item) => {
      const centerObj = {}
      for (const key in dcCenterData[0]) {
        if (item.bk_property_id === key) {
          ;(centerObj['bk_property_name'] = item.bk_property_name),
            (centerObj['bk_property_value'] = dcCenterData[0][key])
          centerObj['bk_group_id'] = item.bk_group_id
          centerObj['bk_property_index'] = item.bk_property_index
        }
      }
      result.push(centerObj)
    })
    arr.push(obj)
    // console.log(result, 'result', arr)
    for (let q = 0; q < arr.length; q++) {
      const test = []
      result.map((item) => {
        if (arr[q].bk_group_id === item.bk_group_id) {
          test.push(item)
          arr[q]['data'] = test
        }
      })
    }
  }
  // console.log(arr, 'arr')
  return arr
}

export const dataProcessEdit = () => {
  const dcCenterC = JSON.parse(JSON.stringify(dcCenter))
  const objectattrC = JSON.parse(JSON.stringify(objectattr))

  let arr = []
  for (let i = 0; i < dcCenterC.length; i++) {
    const result = []
    const obj = {}
    for (let j = 0; j < objectattrC.length; j++) {
      if (dcCenterC[i].bk_group_id === objectattrC[j].bk_property_group) {
        // obj['data'] =
        result.push(objectattrC[j])
        // dcCenterC[i][objectattrC[j].bk_property_id] = ''
        obj[objectattrC[j].bk_property_id] = ''
        dcCenterC[i]['data'] = arrItemSort(result, 'bk_property_index', 0)
        dcCenterC[i]['list'] = obj
      }
    }
    arr = dcCenterC
  }

  return arr
}
