/* eslint-disable @typescript-eslint/explicit-module-boundary-types */
import moment from 'moment'
// 拿到当前的月+日
export const getDate = () => {
  const date = new Date()
  const Separator = '/'
  // const year = date.getFullYear()
  let month = (date.getMonth() + 1).toString()
  let strDate = (date.getDate() + 0).toString()
  if (parseInt(month) >= 1 && parseInt(month) <= 9) {
    month = '0' + month
  }
  if (parseInt(strDate) >= 0 && parseInt(strDate) <= 9) {
    strDate = '0' + strDate
  }
  const data = month + Separator + strDate
  return data
}

// 格式化日期
export const formatDate = (data, format: string) => {
  return moment(data).format(format)
}

/**
 * 邮箱
 * @param {*} s
 */
export function isEmail(s) {
  return /^([a-zA-Z0-9_-])+@([a-zA-Z0-9_-])+((.[a-zA-Z0-9_-]{2,3}){1,2})$/.test(
    s,
  )
}

/**
 * 手机号码
 * @param {*} s
 */
export function isMobile(s) {
  return /^1[0-9]{10}$/.test(s)
}

/**
 * 电话号码
 * @param {*} s
 */
export function isPhone(s) {
  return /^([0-9]{3,4}-)?[0-9]{7,8}$/.test(s)
}

/**
 * URL地址
 * @param {*} s
 */
export function isURL(s) {
  return /^http[s]?:\/\/.*/.test(s)
}

/* 小写字母 */
export function isLowerCase(str) {
  const reg = /^[a-z]+$/
  return reg.test(str)
}

/* 大写字母 */
export function isUpperCase(str) {
  const reg = /^[A-Z]+$/
  return reg.test(str)
}

/* 大小写字母 */
export function isAlphabets(str) {
  const reg = /^[A-Za-z]+$/
  return reg.test(str)
}

/**
 * 判断姓名是否正确
 */
export function isName(name) {
  const regName = /^[\u4e00-\u9fa5]{2,4}$/
  if (!regName.test(name)) return false
  return true
}

/**
 * 判断是否为整数
 */
export function isNum(num, type) {
  let regName = /[^\d.]/g
  if (type === 1) {
    if (!regName.test(num)) return false
  } else if (type === 2) {
    regName = /[^\d]/g
    if (!regName.test(num)) return false
  }
  return true
}

/**
 * 判断是否为小数
 */
export function isNumOrd(num, type) {
  let regName = /[^\d.]/g
  if (type === 1) {
    if (!regName.test(num)) return false
  } else if (type === 2) {
    regName = /[^\d.]/g
    if (!regName.test(num)) return false
  }
  return true
}

/**
 * 判断是否为空
 */
export function isNull(val) {
  if (val instanceof Array) {
    if (val.length === 0) return true
  } else if (val instanceof Object) {
    if (JSON.stringify(val) === '{}') return true
  } else {
    if (
      val === 'null' ||
      val == null ||
      val === 'undefined' ||
      val === undefined ||
      val === ''
    )
      return true
    return false
  }
  return false
}

/**
 * 判断身份证号码
 */
export function isCardId(code) {
  let msg = ''
  const city = {
    11: '北京',
    12: '天津',
    13: '河北',
    14: '山西',
    15: '内蒙古',
    21: '辽宁',
    22: '吉林',
    23: '黑龙江 ',
    31: '上海',
    32: '江苏',
    33: '浙江',
    34: '安徽',
    35: '福建',
    36: '江西',
    37: '山东',
    41: '河南',
    42: '湖北 ',
    43: '湖南',
    44: '广东',
    45: '广西',
    46: '海南',
    50: '重庆',
    51: '四川',
    52: '贵州',
    53: '云南',
    54: '西藏 ',
    61: '陕西',
    62: '甘肃',
    63: '青海',
    64: '宁夏',
    65: '新疆',
    71: '台湾',
    81: '香港',
    82: '澳门',
    91: '国外 ',
  }
  if (!isNull(code)) {
    if (code.length === 18) {
      if (!code || !/(^\d{18}$)|(^\d{17}(\d|X|x)$)/.test(code)) {
        msg = '证件号码格式错误'
        return false
      } else if (!city[code.substr(0, 2)]) {
        msg = '地址编码错误'
        return false
      } else {
        // 18位身份证需要验证最后一位校验位
        code = code.split('')
        // ∑(ai×Wi)(mod 11)
        // 加权因子
        const factor = [7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2]
        // 校验位
        const parity = [1, 0, 'X', 9, 8, 7, 6, 5, 4, 3, 2, 'x']
        let sum = 0
        let ai = 0
        let wi = 0
        for (let i = 0; i < 17; i++) {
          ai = code[i]
          wi = factor[i]
          sum += ai * wi
        }
        const last = parity[sum % 11]
        if (last !== code[17]) {
          msg = '证件号码校验位错误'
          return false
        }
      }
    } else {
      msg = '证件号码长度不为18位'
      return false
    }
  } else {
    msg = '证件号码不能为空'
    return false
  }
  if (msg) {
    console.log(msg)
  }
  return true
}

/**
 * 树形数据转换
 * @param {*} data list数据
 * @param {*} id 主键ID
 * @param {*} pid 上级ID
 * @param childrenKey 子list数据的key
 */
export function treeDataTranslate(
  data,
  id = 'id',
  pid = 'parentId',
  childrenKey = 'children',
) {
  const res = []
  const temp = {}
  for (let i = 0; i < data.length; i++) {
    temp[data[i][id]] = data[i]
  }
  for (let k = 0; k < data.length; k++) {
    if (temp[data[k][pid]] && data[k][id] !== data[k][pid]) {
      if (!temp[data[k][pid]][childrenKey]) {
        temp[data[k][pid]][childrenKey] = []
      }
      if (!temp[data[k][pid]]['_level']) {
        temp[data[k][pid]]['_level'] = 1
      }
      data[k]['_level'] = temp[data[k][pid]]._level + 1
      temp[data[k][pid]][childrenKey].push(data[k])
    } else {
      res.push(data[k])
    }
  }
  return res
}

// 数组去重
export const uniqArray = (array) => {
  const temp = [] //一个新的临时数组
  const newArr = []
  for (let i = 0; i < array.length; i++) {
    if (temp.indexOf(array[i].C_KEYVALUE) == -1) {
      temp.push(array[i].C_KEYVALUE)
      newArr.push(array[i])
    }
  }

  return newArr
}
// filters数据
export const getFilterData = (headersearchlist) => {
  const selectData = headersearchlist
  const filterData = []
  if (selectData) {
    for (let i = 0; i < selectData.length; i++) {
      const value = selectData[i].C_KEYVALUE + '-' + selectData[i].C_CAPTION
      filterData.push({
        text: value,
        value: selectData[i].C_KEYVALUE,
        key: selectData[i].C_KEYVALUE,
      })
    }
  }

  return filterData
}

// 格式化日期
export const getDate1 = (date) => {
  let dateStr = ''
  if (!(date instanceof Object)) {
    const year = date.substring(0, 4)
    const month = date.substring(4, 6)
    const day = date.substring(6, 8)
    dateStr = year + '-' + month + '-' + day
  } else {
    const y = new Date(date).getFullYear().toString()
    const m =
      new Date(date).getMonth() + 1 < 10
        ? '0' + (new Date(date).getMonth() + 1)
        : new Date(date).getMonth() + 1
    const d =
      new Date(date).getDate() < 10
        ? '0' + new Date(date).getDate()
        : new Date(date).getDate()
    dateStr = y + m + d
  }

  return dateStr
}
export const setItem = (key, value) => {
  // 将数组、对象类型的数据转化为 JSON 字符串进行存储
  if (typeof value === 'object') {
    value = JSON.stringify(value)
  }
  window.localStorage.setItem(key, value)
}

/**
 * 获取数据
 */
export const getItem = (key) => {
  const data = window.localStorage.getItem(key)
  try {
    return JSON.parse(data)
  } catch (err) {
    return data
  }
}

/**
 * 删除数据
 */
export const removeItem = (key) => {
  window.localStorage.removeItem(key)
}

/**
 * 删除所有数据
 */
export const removeAllItem = () => {
  window.localStorage.clear()
}

export function getCurrentAuthority() {
  return ['admin']
}

export function check(authority) {
  const current = getCurrentAuthority()
  return current.some((item) => authority.includes(item))
}
export const getMenList = (authList) => {
  const menu = []
  const sourceMap = {}
  authList.forEach((m) => {
    m.children = []
    sourceMap[m.id] = m
    if (m.pid === -1) {
      menu.push(m)
    } else {
      sourceMap[m.pid] && sourceMap[m.pid].children.push(m)
    }
  })
  return menu
}

export const arrItemSort = (arrObj, keyName, type) => {
  //这里如果 直接等于arrObj，相当于只是对对象的引用，改变排序会同时影响原有对象的排序，而通过arrObj.slice(0)，
  //表示把对象复制给另一个对象，两者间互不影响
  const tempArrObj = arrObj.slice(0)
  const compare = function (keyName, type) {
    return function (obj1, obj2) {
      let val1 = obj1[keyName]
      let val2 = obj2[keyName]
      if (!isNaN(Number(val1)) && !isNaN(Number(val2))) {
        val1 = Number(val1)
        val2 = Number(val2)
      }
      //如果值为空的，放在最后
      if (val1 == null && val2 == null) {
        return 0
      } else if (val1 == null && val2 != null) {
        return type == 1 ? -1 : 1
      } else if (val2 == null && val1 != null) {
        return type == 1 ? 1 : -1
      }
      //排序
      if (val1 < val2) {
        return type == 1 ? 1 : -1
      } else if (val1 > val2) {
        return type == 1 ? -1 : 1
      } else {
        return 0
      }
    }
  }
  return tempArrObj.sort(compare(keyName, type))
}
