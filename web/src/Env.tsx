// Copyright (c) 2022-2023 - for information on the respective copyright owner
// see the NOTICE file and/or the repository at
// https://github.com/catenax-ng/product-esc-backbone-code
//
// SPDX-License-Identifier: Apache-2.0

export default class Env{
    private static env:Promise<EnvVars>
    private constructor() { }
    public static getVars():Promise<EnvVars>{
        if(!Env.env){
            Env.env=fetch("/chain/env.json").then(e=>e.json()).catch(console.log)
        }
        return Env.env
    }
}

export interface EnvVars {
    [key: string]: string;
}