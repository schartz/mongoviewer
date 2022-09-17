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
        <a v-on:click="activeTab = 'Users'">Users ({{activeDatabase?.users?.length}})</a>
      </li>
    </ul>
  </div>
  <div class="tab-contents">
    <div v-if="activeTab === 'Collections'" class="content">
      <div v-for="coll in activeDatabase?.collections" class="columns collection-row">
        <div class="column is-10">
          {{coll}}
        </div>
        <div class="column">
          <span class="tag is-primary">Open</span>
        </div>
      </div>
    </div>
    <div v-if="activeTab === 'Functions'" class="content" >
      <div v-for="fn in activeDatabase?.functions" class="columns collection-row">
        <div class="column">
          {{fn._id}}
        </div>
        <div class="column">
          {{fn.value}}
        </div>
      </div>
    </div>
    <div v-if="activeTab === 'Users'" class="content">
      Users content
    </div>
  </div>
</template>

<style scoped lang="scss">
@import "/@src/scss/theme/_variables.scss";
.tab-contents {
  .content {
    padding:1rem
  }
}
.collection-row {
  &:hover {
    background: $simple-bg-hover;
  }

  span.tag {
    cursor:pointer;
    &:hover {
      background: #20344a;
    }
  }
}
</style>