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