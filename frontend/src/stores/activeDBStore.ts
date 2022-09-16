import {defineStore} from "pinia";
import {useLocalStorage} from "@vueuse/core";

export const activeDBStore  = defineStore('activeDBSTore', () => {
    const activeDB = useLocalStorage<ActiveDatabase>('activeDatabase', {
        name: "some database",
        collections: ["coll1", "coll2"],
        functions: ["func1", "func2"],
        users: ["user1", "user2"]
    })

    const updateActiveDB = (dbData: ActiveDatabase) => {
        activeDB.value = dbData
    }

    const getActiveDB = () => {
        return activeDB.value
    }

    return {updateActiveDB, getActiveDB} as const
})