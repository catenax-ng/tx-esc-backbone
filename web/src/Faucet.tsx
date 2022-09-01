import React, {useState} from 'react';
import {Button, Input} from "cx-portal-shared-components";
import Env from "./Env";

function SendFunds(receiver: string, faucetAddress: string) {
    fetch(faucetAddress+"/credit", {
        method: "POST",
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({"denom":"ncaxdemo","address":receiver})
        }
    ).catch(console.log).then((_)=>{})
}
const URL_REGEX = new RegExp('(https?:\\/\\/[A-Za-z-_.]+(:\\d+)?)(\\/.*)?');

function sanitizeUrlToHostWithSchemaAndPort(url: string, openError: (errorMsg: string) => void): string {
    const match = url.match(URL_REGEX)
    if (match) {
        return match[1]
    }else{
        throw "Invalid faucet url: "+ url
    }
}

export default function Faucet() {
    const [receiver, setReceiver] = useState("");
    const faucetAddress=Env.getVars().then(o => o["WEBAPP_FAUCET"]).then(u => sanitizeUrlToHostWithSchemaAndPort(u))
    return (
        <div>
            <Input
                helperText="Put your public address here (starts with 'catenax1')"
                label="Public address"
                onClick={function noRefCheck() {
                }}
                placeholder="enter your public address here"
                onChange={(event) => setReceiver((event.target.value))}
            />
            <Button
                color="primary"
                onClick={()=> { faucetAddress.then(a=>SendFunds( receiver.toString(),a))}}
                onFocusVisible={function noRefCheck() {
                }}
                size="large"
            >Fund me</Button>
        </div>
    );
}