<script setup lang="ts">
import { Codemirror } from 'vue-codemirror'
import {oneDark} from "@codemirror/theme-one-dark";
import {javascript} from "@codemirror/lang-javascript";
import {onMounted, ref, shallowRef} from "vue";
import {useOpenCollectionStore} from "/@src/stores/openMongoCollection";
import {GetCollectionItems} from "../../../wailsjs/go/main/App";

const code = ref('')
const filter = ref(`.find({})
   .projection({})
   .sort({_id:-1})
   .limit(100)`)
const extensions = [javascript(), oneDark]
const collectionName = ref('')
const openCollectionStore = useOpenCollectionStore()

const view = shallowRef()
const handleReady = (payload: any) => {
  view.value = payload.view
}

const log = console.log

const runQuery = () => {
  GetCollectionItems(collectionName.value, filter.value)
      .then(response => {
        console.log("***//*** ", response)
      })
      .catch(err => {
        console.err(err)
      })
}

onMounted(() => {
  collectionName.value = openCollectionStore.getName()
  code.value = `db.${collectionName.value}` + filter.value
})

</script>

<template>
  <div>
    <codemirror
        v-model="code"
        placeholder="Code goes here..."
        :style="{ height: '200px' }"
        :autofocus="true"
        :indent-with-tab="true"
        :tab-size="2"
        :extensions="extensions"
        @ready="handleReady"
    />
  </div>
  <hr class="divider">
  <div>
    <table class="table">
      <thead>
      <tr>
        <th><abbr title="Key">Key</abbr></th>
        <th><abbr title="Value">Value</abbr></th>
        <th><abbr title="Type">Type</abbr></th>
      </tr>
      </thead>
      <tbody>
      <tr>
        <td></td>
        <td></td>
        <td></td>
      </tr>
      </tbody>
    </table>

  </div>
</template>

<style scoped>

</style>