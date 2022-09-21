interface ConnectionData {
    name: string,
    uri: string,
}

interface DbInfo {
    Name: string,
    SizeOnDisk: number,
    Empty: boolean
}

interface ActiveDatabase {
    name: string,
    collections: Array<string>,
    functions: Array<DBFunction>,
    users: any
}

interface DBFunction {
    _id: string
    value: string
}

interface ConnectionListTreeRoot {
    text: string,
    state: ConnectionListTreeNodeState,
    id: number,
    checkable: boolean,
    nodes: Array<ConnectionListTreeNode>

}

interface ConnectionListTreeNodeState {
    checked: boolean,
    selected: boolean,
    expanded: boolean
}

interface ConnectionListTreeNode {
    text: string,
    state: ConnectionListTreeNodeState,
    id: number,
    nodes: Array<ConnectionListTreeNode>
}