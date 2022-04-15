// 保存数据到sessionStorage
export function setSession(key: string, value: string) {
  sessionStorage.setItem(key, value)
}

// 从sessionStorage获取数据
export function getSession(key: string) {
  return sessionStorage.getItem(key)
}

// 从sessionStorage删除保存的数据
export function delSession(key: string) {
  sessionStorage.removeItem(key)
}

// 从sessionStorage删除所有保存的数据
export function clearSession() {
  sessionStorage.clear()
}

export function setLocal(key: string, value: string) {
  localStorage.setItem(key, value)
}

export function getLocal(key: string) {
  return localStorage.getItem(key)
}

export function delLocal(key: string) {
  localStorage.removeItem(key)
}

export function clearLocal() {
  localStorage.clear()
}

export function setCookie(name: string, value: string) {
  const Days = 30
  const exp = new Date()
  exp.setTime(exp.getTime() + Days * 24 * 60 * 60 * 1000)
  // eslint-disable-next-line @typescript-eslint/ban-ts-comment
  // @ts-ignore
  document.cookie = name + '=' + escape(value) + ';expires=' + exp.toGMTString()
}

export function getCookie(name: string) {
  const reg = new RegExp('(^| )' + name + '=([^;]*)(;|$)')
  const arr = document.cookie.match(reg)
  if (arr && arr.length > 0) {
    return unescape(arr[2])
  } else {
    return null
  }
}

export function delCookie(name: string) {
  const exp = new Date()
  exp.setTime(exp.getTime() - 1)
  const cval = getCookie(name)
  // eslint-disable-next-line @typescript-eslint/ban-ts-comment
  // @ts-ignore
  if (cval != null) {
    document.cookie = name + '=' + cval + ';expires=' + exp.toGMTString()
  }
}

export function clearAllCookie() {
  const keys = document.cookie.match(/[^ =;]+(?==)/g)
  if (keys) {
    for (let i = keys.length; i--;) {
      document.cookie =
        keys[i] + '=0;path=/;expires=' + new Date(0).toUTCString() // 清除当前域名下的
      document.cookie =
        keys[i] +
        '=0;path=/;domain=' +
        document.domain +
        ';expires=' +
        new Date(0).toUTCString() // 清除当前域名下的
      document.cookie =
        keys[i] +
        '=0;path=/;domain=ratingdog.cn;expires=' +
        new Date(0).toUTCString() // 清除一级域名下的或指定的
    }
  }
}
