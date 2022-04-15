import { markRaw, reactive } from 'vue'
import { CirclePlusFilled } from '@element-plus/icons-vue'

export const buttonData = reactive({
  buttonIn: {
    name: '导入',
    className: 'roundBackCal',
    round: true,
  },
  buttonOut: {
    name: '导出',
    className: 'roundBackCal',
    round: true,
  },
  buttonSubmit: {
    name: '提交',
    className: 'roundBackCal',
    round: true,
  },
  buttonCancel: {
    name: '取消',
    className: 'roundBackCal',
    round: true,
  },
  buttonNew: {
    name: '新建关联',
    className: 'roundWhiteIcon',
    round: true,
    icon: markRaw(CirclePlusFilled),
  },
  buttonInPart: {
    name: '导入',
    className: 'roundWhiteIcon',
    round: true,
  },
  buttonOutPart: {
    name: '导出',
    className: 'roundWhiteIcon',
    round: true,
  },
})


export const uploadData = reactive({
  title:'导入',
  isShowUpLoad:false,
  upLoad(){
    uploadData.isShowUpLoad = true;
  },
  downLoad(){
   console.log('导出')
  },
  btnCancel(){
    uploadData.isShowUpLoad = false;
  },
  btnSure(){
    uploadData.isShowUpLoad = false;
  }

})