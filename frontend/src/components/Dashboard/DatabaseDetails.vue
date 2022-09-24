<script setup lang="ts">
import {onMounted, ref, shallowRef} from "vue";
import {useActiveDBStore} from "/@src/stores/useActiveDBStore";
import CollectionDetails from "/@src/components/partials/CollectionDetails.vue";
import {useOpenCollectionStore} from "/@src/stores/openMongoCollection";



const activeDatabaseStore = useActiveDBStore()
const activeTab = ref<string>('Collections')
const activeDatabase = ref<ActiveDatabase>()
const selectedCollection = ref("")
const openCollectionStore = useOpenCollectionStore()

const init = async() => {
  activeDatabase.value = activeDatabaseStore.getActiveDB()
}

const openCollection = (collectionsName: string) => {
  selectedCollection.value = collectionsName
  openCollectionStore.updateName(collectionsName)
}

onMounted(async() => {
  await init()
})

</script>

<template>
  <div class="tabs is-centered">
    <ul>
      <li v-bind:class="{ 'is-active': activeTab === 'Collections' }">
        <a v-on:click="activeTab = 'Collections'">Collections ({{activeDatabase?.collections?.length || 0}})</a>
      </li>
      <li v-bind:class="{ 'is-active': activeTab === 'Functions' }">
        <a v-on:click="activeTab = 'Functions'">Functions ({{activeDatabase?.functions?.length || 0}})</a>
      </li>
      <li v-bind:class="{ 'is-active': activeTab === 'Users' }">
        <a v-on:click="activeTab = 'Users'">Users ({{activeDatabase?.users.users?.length || 0}})</a>
      </li>
    </ul>
  </div>
  <div class="tab-contents">
    <div v-if="activeTab === 'Collections'" class="content">
      <div v-if="selectedCollection !== ''">
        <div class="select is-fullwidth">
          <select v-model="selectedCollection">
            <option value="">Select a collection to open</option>
            <option v-bind:value="coll" v-for="coll in activeDatabase?.collections">{{coll}}</option>
          </select>
        </div>
        <hr class="divider">
        <collection-details></collection-details>
      </div>
      <div v-else>
        <div v-for="coll in activeDatabase?.collections" class="columns collection-row">
          <div class="column is-10">
            {{coll}}
          </div>
          <div class="column">
            <span class="tag is-primary" @click="openCollection(coll)">Open</span>
          </div>
        </div>
      </div>
    </div>
    <div v-if="activeTab === 'Functions'" class="content" >
      <div v-for="fn in activeDatabase?.functions" class="columns collection-row">
        <div class="column is-2">
          {{fn[0]['Value']}}
        </div>
        <div class="column has-text-centered">
          <span style="font-family: 'monospace'">
            {{fn[1]['Value']}}
          </span>
        </div>
        <div class="column is-1">
          <span class="tag is-warning">Edit</span>
        </div>
      </div>
    </div>
    <div v-if="activeTab === 'Users'" class="content">

      <div v-for="userObj in activeDatabase?.users.users" class="box">
        <div class="columns collection-row">
          <div class="column">
            _id
          </div>
          <div class="column">
          <span style="font-family: 'monospace'">
            {{userObj._id}}
          </span>
          </div>
        </div>
        <div class="columns collection-row">
          <div class="column">
            username
          </div>
          <div class="column">
          <span style="font-family: 'monospace'">
            {{userObj.user}}
          </span>
          </div>
        </div>
        <div class="columns collection-row">
          <div class="column">
            roles
          </div>
          <div class="column">
          <span style="font-family: 'monospace'">
            {{userObj.roles}}
          </span>
          </div>
        </div>
        <div class="columns collection-row">
          <div class="column">
            mechanisms
          </div>
          <div class="column">
          <span style="font-family: 'monospace'">
            {{userObj.mechanisms}}
          </span>
          </div>
        </div>
      </div>

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