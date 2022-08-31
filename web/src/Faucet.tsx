import React, {useState} from 'react';
import {Button, Input} from "cx-portal-shared-components";


export default function Faucet() {
    const [receiver, setReceiver] = useState("");
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
                onClick={function noRefCheck() {
                    console.log("Public address", {receiver})
                }}
                onFocusVisible={function noRefCheck() {
                }}
                size="large"
            >Fund me</Button>
        </div>
    );
}