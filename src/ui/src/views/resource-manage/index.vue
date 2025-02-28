<template>
  <div class="classify-layout clearfix" v-bkloading="{ isLoading: $loading('getObjectCommonInstanceCount') }">
    <div class="classify-filter">
      <bk-input class="filter-input"
        clearable
        :placeholder="$t('请输入xx', { name: $t('关键字') })"
        right-icon="icon-search"
        v-model.trim="filter">
      </bk-input>
    </div>
    <div v-show="!isEmpty">
      <div class="classify-waterfall fl"
        v-for="col in classifyColumns.length"
        :key="col">
        <cmdb-classify-panel
          v-for="classify in classifyColumns[col - 1]"
          :key="classify['bk_classification_id']"
          :classify="classify"
          :collection="collection"
          :instance-count="instanceCount">
        </cmdb-classify-panel>
      </div>
    </div>
    <no-search-results v-if="isEmpty && !globalLoading" :text="$t('搜不到相关资源')" />
  </div>
</template>

<script>
  import { mapGetters } from 'vuex'
  import {
    MENU_RESOURCE_COLLECTION,
    MENU_RESOURCE_HOST_COLLECTION,
    MENU_RESOURCE_BUSINESS_COLLECTION
  } from '@/dictionary/menu-symbol'
  import noSearchResults from '@/views/status/no-search-results.vue'
  import cmdbClassifyPanel from './children/classify-panel'
  import debounce from 'lodash.debounce'
  export default {
    components: {
      cmdbClassifyPanel,
      noSearchResults
    },
    data() {
      return {
        filter: '',
        debounceFilter: null,
        matchedModels: null,
        instanceCount: []
      }
    },
    computed: {
      ...mapGetters(['globalLoading']),
      ...mapGetters('objectModelClassify', ['classifications', 'models']),
      ...mapGetters('userCustom', ['usercustom']),
      collection() {
        const isHostCollected = this.usercustom[MENU_RESOURCE_HOST_COLLECTION] === undefined
          ? true
          : this.usercustom[MENU_RESOURCE_HOST_COLLECTION]
        const isBusinessCollected = this.usercustom[MENU_RESOURCE_BUSINESS_COLLECTION] === undefined
          ? true
          : this.usercustom[MENU_RESOURCE_BUSINESS_COLLECTION]
        const collection = [...(this.usercustom[MENU_RESOURCE_COLLECTION] || [])]
        if (isHostCollected) {
          collection.push('host')
        }
        if (isBusinessCollected) {
          collection.push('biz')
        }
        return collection.filter(modelId => this.models.some(model => model.bk_obj_id === modelId))
      },
      filteredClassifications() {
        const result = []
        const filterClassify = ['bk_biz_topo']
        this.classifications.forEach((classification) => {
          if (!filterClassify.includes(classification.bk_classification_id)) {
            const models = classification.bk_objects.filter((model) => {
              const isInvisible = model.bk_ishidden
              const isPaused = model.bk_ispaused
              const isMatched = this.matchedModels ? this.matchedModels.includes(model.bk_obj_id) : true
              return !isInvisible && !isPaused && isMatched
            })
            if (models.length) {
              result.push({
                ...classification,
                bk_objects: models
              })
            }
          }
        })
        return result
      },
      classifyColumns() {
        const colHeight = [0, 0, 0, 0]
        const classifyColumns = [[], [], [], []]
        this.filteredClassifications.forEach((classify) => {
          const minColHeight = Math.min(...colHeight)
          const rowIndex = colHeight.indexOf(minColHeight)
          classifyColumns[rowIndex].push(classify)
          colHeight[rowIndex] = colHeight[rowIndex] + this.calcWaterfallHeight(classify)
        })
        return classifyColumns
      },
      isEmpty() {
        return this.classifyColumns.every(column => !column.length)
      }
    },
    watch: {
      filter() {
        this.debounceFilter()
      }
    },
    created() {
      this.debounceFilter = debounce(this.filterModel, 300)
      this.getInstanceCount()
    },
    methods: {
      filterModel() {
        if (this.filter) {
          const models = this.models.filter(model => model.bk_obj_name.indexOf(this.filter) > -1)
          this.matchedModels = models.map(model => model.bk_obj_id)
        } else {
          this.matchedModels = null
        }
      },
      async getInstanceCount() {
        try {
          this.instanceCount = await this.$store.dispatch('objectCommonInst/getInstanceCount', {
            config: {
              requestId: 'getObjectCommonInstanceCount',
              globalError: false
            }
          })
        } catch (e) {
          console.error(e)
          this.instanceCount = []
          this.$route.meta.view = 'error'
        }
      },
      calcWaterfallHeight(classify) {
        // 46px 分类高度
        // 16px 模型列表padding
        // 36 模型高度
        return 46 + 16 + (classify.bk_objects.length * 36)
      }
    }
  }
</script>

<style lang="scss" scoped>
    .classify-layout{
        padding: 15px 20px 20px;
    }
    .classify-filter {
        padding: 0 20px 20px 0;
        .filter-input {
            width: 240px;
        }
    }
    .classify-waterfall{
        width: calc((100% - 80px) / 4);
        margin: 0 0 0 20px;
        &:first-child{
            margin: 0;
        }
    }
</style>
