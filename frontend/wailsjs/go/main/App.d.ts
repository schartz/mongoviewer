// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {mongo} from '../models';

export function AddConnection(arg1:string):Promise<boolean>;

export function ConnectToDBServer(arg1:string):Promise<Array<mongo.DatabaseSpecification>>;

export function ConnectionList():Promise<Array<{[key: string]: string}>>;

export function Greet(arg1:string):Promise<string>;

export function TestConnection(arg1:string):Promise<string>;
