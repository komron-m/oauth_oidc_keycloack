import React, {useEffect, useState} from 'react';
import axios from "./axios";

interface Hero {
    id: string
    name: string
    description: string
}

const host = "http://localhost:4000"

function Table() {
    const [heroes, setHeroes] = useState<Hero[]>()
    const [name, setName] = useState("")
    const [description, setDescription] = useState("")

    useEffect(() => {
        getAll()
    }, [])

    const getAll = () => {
        axios.get(`${host}/get_all`).then((resp) => {
            setHeroes(resp.data)
        }).catch((err) => {
            console.log(err)
        })
    }

    const deleteHero = (id: string) => {
        axios.post(`${host}/delete`, {id}).then((resp) => {
            getAll()
        }).catch((err) => {
            console.log(err)
        })
    }

    const onSave = (e: any) => {
        if (name === "" || description === "") {
            return
        }
        axios.post(`${host}/create`, {
            name, description
        }).then((resp) => {
            getAll()
        }).catch((err) => {
            console.log(err)
        }).finally(() => {
            setName("")
            setDescription("")
        })
    }

    return (
        <>
            <form onSubmit={(e) => e.preventDefault()}>
                <div className={"p-2"}>
                    <label htmlFor={"name"}>Name</label>&nbsp;
                    <input
                        value={name}
                        onChange={(e) => setName(e.target.value)}
                        className={"border"} type={"text"} id={"name"} required={true}/>
                </div>
                <div className={"p-2"}>
                    <label htmlFor={"description"}>Description</label>&nbsp;
                    <input
                        value={description}
                        onChange={(e) => setDescription(e.target.value)}
                        className={"border"} type={"text"} id={"description"} required={true}/>
                </div>

                <div className={"p-2"}>
                    <button
                        onClick={onSave}
                        className={"rounded-md text-white bg-green-600 py-1 px-2"}
                        type={"submit"}>Save
                    </button>
                </div>
            </form>

            <table className="border-collapse border border-slate-500">
                <thead>
                <tr>
                    <th className="border p-2">ID</th>
                    <th className="border p-2">NAME</th>
                    <th className="border p-2">DESCRIPTION</th>
                    <th className="border p-2">-</th>
                </tr>
                </thead>
                <tbody>
                {
                    heroes?.map((h) => {
                        return (
                            <tr key={h.id}>
                                <td className="border p-2">
                                    {h.id.substring(0, 8)}
                                </td>
                                <td className="border p-2">{h.name}</td>
                                <td className="border p-2">
                                    <i>{h.description}</i>
                                </td>
                                <td className="border p-2">
                                    <button
                                        onClick={(e) => deleteHero(h.id)}
                                        className={"rounded-md text-white bg-red-600 px-1"}>
                                        {"X"}
                                    </button>
                                </td>
                            </tr>
                        )
                    })
                }
                </tbody>
            </table>
        </>
    )
}

function App() {
    return (
        <>
            <div className="flex">
                <div className="flex-1 w-1/3"/>
                <div className="flex-1 w-1/3 p-10">
                    <Table/>
                </div>
                <div className="flex-1 w-1/3"/>
            </div>
        </>
    );
}

export default App;
