<script setup lang="ts">
import {onMounted, ref} from "vue";
import {activeDBStore} from "/@src/stores/activeDBStore";

const activeDatabaseStore = activeDBStore()
const activeTab = ref<string>('Collections')
const activeDatabase = ref<ActiveDatabase>()

const init = async() => {
  activeDatabase.value = activeDatabaseStore.getActiveDB()
}

onMounted(async() => {
  await init()
})

</script>

<template>
  <h1 class="has-text-centered mb-5 mt-4 is-size-5" style="background: #1d2122">
    Active Database : {{activeDatabase?.name}}
  </h1>
  <div class="tabs is-centered">
    <ul>
      <li v-bind:class="{ 'is-active': activeTab === 'Collections' }">
        <a v-on:click="activeTab = 'Collections'">Collections ({{activeDatabase?.collections.length}})</a>
      </li>
      <li v-bind:class="{ 'is-active': activeTab === 'Functions' }">
        <a v-on:click="activeTab = 'Functions'">Functions ({{activeDatabase?.functions.length}})</a>
      </li>
      <li v-bind:class="{ 'is-active': activeTab === 'Users' }">
        <a v-on:click="activeTab = 'Users'">Users ({{activeDatabase?.users.length}})</a>
      </li>
    </ul>
  </div>
  <div class="tab-contents">
    <div v-if="activeTab === 'Collections'" class="content">
      <p v-for="coll in activeDatabase?.collections">
        {{coll}} <button class="button is-small is-link">Open</button>
      </p>
    </div>
    <div v-if="activeTab === 'Functions'" class="content" >
      Functions content
    </div>
    <div v-if="activeTab === 'Users'" class="content">
      Users content
    </div>
  </div>
</template>

<style scoped>

</style>