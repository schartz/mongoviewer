<script lang="ts" setup>
import {onMounted, reactive, ref} from 'vue'
import {useRouter} from "vue-router";
import {useDatabaseListStore} from "/@src/stores/databaseList";
import {useCurrentConnectionStore} from "/@src/stores/useCurrentConnectionStore";
import DatabaseDetails from "/@src/components/Dashboard/DatabaseDetails.vue";
import {GetDBDetails} from "../../wailsjs/go/main/App";
import {useActiveDBStore} from "/@src/stores/useActiveDBStore";
import { Splitpanes, Pane } from 'splitpanes'
import 'splitpanes/dist/splitpanes.css'

const storeDBList = useDatabaseListStore()
const storeActiveConnection = useCurrentConnectionStore()
const storeActiveDb = useActiveDBStore()

const activeDbName = ref<string>()
const router = useRouter()
const dbList = ref<Array<DbInfo>>([])
const activeConnection = ref<ConnectionData>()

const isReady = ref<boolean>(false)


const selectDb = async(dbName: string) => {
  isReady.value = false
  GetDBDetails(dbName).then(response => {
    activeDbName.value = dbName
    console.log('*********')
    console.log(response)
    storeActiveDb.updateActiveDB({
      name: dbName,
      collections: response.collections,
      functions: response.dbFunctions,
      users: response.users || []
    })
    isReady.value = true
  }).catch(err => {
    console.error(err)
  })
}


const init = async() => {
  dbList.value = storeDBList.getList()
  activeConnection.value = storeActiveConnection.getActiveConnection()
}

onMounted(async () => {
  await init()
})

</script>

<template>



  <splitpanes style="height: calc(100vh - 2rem)">
    <pane max-size="30" min-size="10" size="15">
      <aside class="menu p-5">
        <p class="menu-label">
          Databases in <i>{{activeConnection?.name}}</i>
        </p>
        <ul class="menu-list">
          <li v-for="db in dbList" :class="{'is-active': activeDbName === db.Name}">
            <a @click="selectDb(db.Name)">{{db.Name}}</a>
          </li>
        </ul>
      </aside>
    </pane>
    <pane>
      <database-details v-if="isReady"></database-details>
    </pane>
  </splitpanes>

<!--  <div class="columns">
    <div class="column is-3 menu-container">
      <aside class="menu p-5">
        <p class="menu-label">
          Databases in <i>{{activeConnection?.name}}</i>
        </p>
        <ul class="menu-list">
          <li v-for="db in dbList" :class="{'is-active': activeDbName === db.Name}">
            <a @click="selectDb(db.Name)">{{db.Name}}</a>
          </li>
        </ul>
      </aside>
    </div>
    <div class="column">
      <database-details v-if="isReady"></database-details>

    </div>
  </div>-->
  <div id="status-bar">
    <p>
      I am status bar
    </p>
  </div>
</template>

<style scoped lang="scss">
@import "/@src/scss/theme/_variables.scss";
/*.columns{
  margin-bottom: 0;
}*/
.menu-container {
  border-right: 2px solid black;
  /*min-height: calc(100vh - 1.3rem);*/
}
li.is-active, li.is-active:hover {
  background-color: $primary;
}
#status-bar {
  background-color: black;
  height: 2rem;
  padding: 5px;
   p {
      font-size: 12px;
   }
}


</style>
