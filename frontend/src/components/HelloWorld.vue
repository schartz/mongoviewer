<template>
  <section class="hero is-dark is-fullheight has-text-centered">
    <div class="hero-body">
      <div>
        <p class="title mb-6">
          MongoViewer
        </p>
        <p class="subtitle">

        </p>

        <div class="table-wrapper mb-6">
          <table class="table is-striped is-narrow is-fullwidth" v-if="connectionList.length > 0">
            <thead>
            <th align="center">
              Connection
            </th>
            <th>
              URI
            </th>
            <th>
              Action
            </th>
            </thead>

            <tbody>
            <tr v-for="conn in connectionList">
              <td>
                {{conn.name}}
              </td>
              <td>
                {{conn.uri}}
              </td>
              <td>
                <button class="button is-link is-small" @click="connectToDB(conn)">Connect</button>
              </td>
            </tr>
            </tbody>
          </table>
          <h1 v-else>
            No saved connections
          </h1>
        </div>

        <p class="subtitle mt-6">
          <button class="button  is-warning" @click="openAddConnectionModal = true">
            + Add a new connection
          </button>
        </p>
      </div>
    </div>
  </section>


  <div class="modal" :class="{'is-active': openAddConnectionModal}">
    <div class="modal-background"></div>
    <div class="modal-content">
      <div class="box">
        <h5 class="is-size-5 mb-4">
          Please fill in the connection information
        </h5>
        <div class="field">
          <label class="label">Connection name</label>
          <div class="control">
            <input v-model="newConnectionData.name" class="input" type="text" placeholder="Text input">
          </div>
        </div>

        <div class="field">
          <label class="label">Connection URI</label>
          <div class="control">
            <input v-model="newConnectionData.uri" class="input" type="text" placeholder="Text input">
          </div>
        </div>

        <hr class="divider">

        <div class="field is-grouped">
          <div class="control">
            <button class="button is-link" @click="addConnection()">Submit</button>
          </div>
          <div class="control">
            <button class="button is-primary" @click="testConnection">Test Connection</button>
          </div>
          <div class="control">
            <button class="button is-link is-light" @click="openAddConnectionModal = false">Cancel</button>
          </div>
        </div>
        <button class="modal-close is-large" aria-label="close" @click="openAddConnectionModal = false"></button>
      </div>
    </div>

  </div>

</template>

<script setup lang="ts">
// This template is using Vue 3 <script setup> SFCs
// Check out https://v3.vuejs.org/api/sfc-script-setup.html#sfc-script-setup
import {onMounted, ref} from 'vue'
import { Notyf } from 'notyf';
import { useRouter } from 'vue-router';
import {AddConnection, ConnectionList, ConnectToDBServer, TestConnection} from "../../wailsjs/go/main/App";
import {databaseListStore} from "/@src/stores/databaseList";
import {currentConnectionStore} from "/@src/stores/currentConnectionStore";

const dbStore = databaseListStore()
const activeConnectionStore = currentConnectionStore()
const notyf = new Notyf({duration: 5000, position:{x:'center',y:'top'}});
const connectionList = ref<{ [key: string]: string; }[]>([])
const openAddConnectionModal = ref(false)
const router = useRouter();
const newConnectionData = ref<ConnectionData>({
  name: 'default',
  uri: 'mongodb://localhost:27017'
})

const fetchConnectionList = () => {
  ConnectionList().then(response => {
    if (response && response.length > 0){
      connectionList.value = response
    }
  })
}

const testConnection = () => {
  TestConnection(newConnectionData.value.uri).then(response => {
    if (response !== 'yes'){
      notyf.error(response);
    } else {
      notyf.success('Connection Successfull');
    }
  })
}

const addConnection = () => {
  AddConnection(JSON.stringify(newConnectionData.value)).then(response => {
    console.log(response)
    if (!response){
      notyf.error('Error saving connection');
    } else {
      notyf.success('Add connection');
      fetchConnectionList()
      openAddConnectionModal.value = false
    }
  })
}

const connectToDB = (connection: ConnectionData) => {
  ConnectToDBServer(connection.uri).then(async response => {
    console.log(response)
    activeConnectionStore.updateActiveConnection(connection)
    dbStore.updateList(response)
    await router.push({name: 'dblist'});

  }).catch(err => {
    console.error(err)
  })
}



onMounted(() => {
  fetchConnectionList()
  console.log(router)
})
</script>

<style lang="scss">
.hero {
  .hero-body {
    justify-content: center;
    align-items: flex-start !important;

    .table-wrapper{
      min-width: 600px;

      th {
        text-align: center;
      }
    }
  }
}
</style>
