import {defineStore} from "pinia";
import {useLocalStorage} from "@vueuse/core";

export const currentConnectionStore  = defineStore('currentConnection', () => {
    const activeConnection = useLocalStorage<ConnectionData>('activeConnection', {
        name: "default",
        uri: "mongodb://localhost:27017"

    })

    const updateActiveConnection = (conn: ConnectionData) => {
        activeConnection.value = conn
    }

    const getActiveConnection = () => {
        return activeConnection.value
    }

    return {updateActiveConnection, getActiveConnection} as const
})