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

export default function Faucet() {
    const [receiver, setReceiver] = useState("");
    const faucetAddress=Env.getVars().then(o=>o["WEBAPP_FAUCET"])
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