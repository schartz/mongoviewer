import {defineStore} from "pinia";
import {ref} from "vue";
import {useLocalStorage} from "@vueuse/core";

export const useOpenCollectionStore  = defineStore('openMongoCollection', () => {
    const name = useLocalStorage<string>('collectioName', '')

    const updateName = (n: string) => {
        name.value = n
    }

    const getName = () => {
        return name.value
    }

    return {updateName, getName} as const
})