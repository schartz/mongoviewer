<script lang="ts" setup>
import {onMounted, reactive, ref} from 'vue'
import {useRouter} from "vue-router";
import {databaseListStore} from "/@src/stores/databaseList";
import {currentConnectionStore} from "/@src/stores/currentConnectionStore";
import DatabaseDetails from "/@src/components/Dashboard/DatabaseDetails.vue";

const dbListStore = databaseListStore()
const activeConnectionStore = currentConnectionStore()
const router = useRouter()
const dbList = ref<Array<DbInfo>>([])
const activeConnection = ref<ConnectionData>()


const init = async() => {
  dbList.value = dbListStore.getList()
  activeConnection.value = activeConnectionStore.getActiveConnection()
}

onMounted(async () => {
  await init()
})

</script>

<template>

  <div class="columns">
    <div class="column is-3 menu-container">
      <aside class="menu p-5">
        <p class="menu-label">
          Databases in <i>{{activeConnection?.name}}</i>
        </p>
        <ul class="menu-list">
          <li v-for="db in dbList">
            <a>{{db.Name}}</a>
          </li>
        </ul>
      </aside>
    </div>
    <div class="column">
      <database-details></database-details>
    </div>
  </div>
</template>

<style scoped lang="scss">
.menu-container {
  border-right: 2px solid black;
}
</style>
