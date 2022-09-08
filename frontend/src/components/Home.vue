<template>
  <section class="hero is-danger is-fullheight">
    <div class="hero-body">
      <div class="">
        <p class="title">
          MongoViewer
        </p>
        <p class="subtitle">
          Fullheight subtitle
        </p>

        <p class="subtitle" @click="openAddConnectionModal = true">
          + Add a new connection
        </p>

        <div>
          <table class="table">
            <thead>
            <tr>
              Connection
            </tr>
            <tr>
              Action
            </tr>
            </thead>

            <tbody>
            <tr v-for="conn in result">
              <td>
                {{conn.conn}}
              </td>
              <td>
                <button class="button is-primary">Connect</button>
              </td>
            </tr>
            </tbody>
          </table>
        </div>
      </div>
    </div>
  </section>


  <div class="modal" :class="{'is-active': openAddConnectionModal}">
    <div class="modal-background"></div>
    <div class="modal-content">
      <div class="box">

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
            <button class="button is-link">Submit</button>
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

interface ConnectionData {
  name: string,
  uri: string,
}
const notyf = new Notyf({duration: 5000});
const result = ref({})
const openAddConnectionModal = ref(false)
const newConnectionData = ref<ConnectionData>({
  name: '',
  uri: ''
})

const fetchConnectionList = () => {
  window.go.main.App.ConnectionList().then(response => {
    result.value = JSON.parse(response)
    console.log(response)
    console.log(result.value)
  })
}

const testConnection = () => {

  window.go.main.App.TestConnection(newConnectionData.value.uri).then(response => {
    if (response !== 'yes'){
      notyf.error(response);
    } else {
      notyf.success('Connection Successfull');
    }
  })
}



onMounted(() => {
  fetchConnectionList()
})
</script>

<style lang="scss">

</style>
