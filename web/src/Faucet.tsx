// Copyright (c) 2022-2023 Contributors to the Eclipse Foundation
//
// See the NOTICE file(s) distributed with this work for additional
// information regarding copyright ownership.
//
// This program and the accompanying materials are made available under the
// terms of the Apache License, Version 2.0 which is available at
// https://www.apache.org/licenses/LICENSE-2.0.
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.
//
// SPDX-License-Identifier: Apache-2.0

import React, {useState} from 'react';
import {Button, Input, PageNotifications} from "cx-portal-shared-components";
import Env from "./Env";

function SendFunds(receiver: string, faucetAddress: string, openError: (errorMsg: string) => void, openInfo: (errorMsg: string) => void) {
    const faucetEndpoint=faucetAddress + "/credit"
    fetch(faucetEndpoint, {
            method: "POST",
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({"denom": "ncaxdemo", "address": receiver})
        }
    ).then((response: Response) => {
        if (!response.ok) {
                      openError(response.statusText)

        } else {
            openInfo("Funds were sent. Please check your balance after a moment.")
        }
    }).catch(e => {
        openError(e.toString())
    })
}
const URL_REGEX = new RegExp('(https?:\\/\\/[A-Za-z-_.]+(:\\d+)?)(\\/.*)?');

function sanitizeUrlToHostWithSchemaAndPort(url: string, openError: (errorMsg: string) => void): string {
    const match = url.match(URL_REGEX)
    if (match) {
        return match[1]
    } else {
        openError("Invalid faucet url: " + url)
        return ""
    }
}


function showErrorFunc(
    setNotifyShow: (value: (((prevState: boolean) => boolean) | boolean)) => void,
    setNotifyMsg: (value: (((prevState: string) => string) | string)
    ) => void,
): (errorMsg: string) => void {
    return function (errorMsg: string) {
        setNotifyMsg(errorMsg)
        setNotifyShow(true)
    }
}

export default function Faucet() {
    const [receiver, setReceiver] = useState("");
    const [notifyMsg, setNotifyMsg] = useState("")
    const [notifyShow, setNotifyShow] = useState(false)
    const openError = showErrorFunc(setNotifyShow, setNotifyMsg)
    // TODO make a popup on this?
    const openInfo = (msg: string) => console.log(msg)
    const faucetAddress = Env.getVars().then(o => o["WEBAPP_FAUCET"]).then(u => sanitizeUrlToHostWithSchemaAndPort(u, openError))

    return (
        <div>
            <PageNotifications
                // contactLinks="https://portal.dev.demo.catena-x.net/"
                // contactText="Contact"
                description={notifyMsg}
                onCloseNotification={() => {
                    setNotifyShow(false)
                }}
                open={notifyShow}
                severity="error"
                showIcon
                title="Faucet error"
            />
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
                onClick={() => {
                    faucetAddress.then(a => SendFunds(receiver.toString(), a, openError, openInfo)).catch(e => openError(e.toString()))
                }}
                onFocusVisible={function noRefCheck() {
                }}
                size="large"
            >Fund me</Button>
        </div>
    );
}