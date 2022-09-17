import {defineStore} from "pinia";
import {useLocalStorage} from "@vueuse/core";

export const activeDBStore  = defineStore('activeDBSTore', () => {
    const activeDB = useLocalStorage<ActiveDatabase>('activeDatabase', {
        name: "No database selected",
        collections: [],
        functions: [],
        users: []
    })

    const updateActiveDB = (dbData: ActiveDatabase) => {
        activeDB.value = dbData
    }

    const getActiveDB = () => {
        return activeDB.value
    }

    return {activeDB, updateActiveDB, getActiveDB}
})