import {defineStore} from "pinia";
import {useLocalStorage} from "@vueuse/core";

export const databaseListStore  = defineStore('databaseListStore', () => {
    const dbs = useLocalStorage<Array<DbInfo>>('dbs', [])

    const updateList = (lst: Array<DbInfo>) => {
        dbs.value = lst
    }

    const getList = () => {
        return dbs.value
    }

    return {updateList, getList} as const
})