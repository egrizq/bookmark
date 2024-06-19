'use client'

import { useEffect, useState } from "react";
import PopupLogin from "./popupLogin";
import PopupRegister from "./popupRegister";
import TweetPage, { isTwitterUrl } from "./tweet";
import { Spotify } from "react-spotify-embed";
import ButtonStatus from "./buttonSubmit";
import Title from "./ui/mainTitle";
import ErrorLink from "./ui/errorURL";
import { typeResponse } from "../auth/definitions";

const isSpotifyUrl = (url: string) => {
    return url.includes('open.spotify.com')
};

export default function Dashboard() {
    const [sosmed, setSosmed] = useState('tw')
    const [url, setUrl] = useState('')
    const [category, setCategory] = useState('#')
    const [returnedSosmed, setReturnedSosmed] = useState<JSX.Element>()

    const [loginPopUp, setIsPopupLogin] = useState(false);
    const [registerPopUp, setIsPopupRegister] = useState(false);

    const togglePopupLogin = () => {
        setIsPopupLogin(!loginPopUp);
    };

    const togglePopupRegister = () => {
        setIsPopupRegister(!registerPopUp);
        setIsPopupLogin(false);
    };

    const closeRegister = () => {
        setIsPopupRegister(!registerPopUp);
    }

    useEffect(() => {
        if (sosmed === 'tw' && isTwitterUrl({ url })) {
            let embedTW = <TweetPage url={url} />
            setReturnedSosmed(embedTW)

        } else if (sosmed === 'sp' && isSpotifyUrl(url)) {
            let embedSP = <Spotify link={url} />
            setReturnedSosmed(embedSP)

        } else {
            let err = <ErrorLink />
            setReturnedSosmed(err)
        }

    }, [url, sosmed])

    const onSubmit = (e: { preventDefault: () => void; }) => {
        e.preventDefault();

        const responseData: typeResponse = {
            id: "egrizq",
            social: sosmed,
            url: url,
            kategory: category
        }

        let format = JSON.stringify(responseData)
        alert(format)
    }

    return (
        <>
            <nav className="border-b border-zinc-200 mx-auto w-full p-3">
                <div className="flex mx-auto justify-between w-9/12">
                    <p className="text-xl font-bold py-1 text-zinc-800">Bookmark-ku</p>
                    <button onClick={togglePopupLogin}
                        className="text-lg text-white font-medium py-1 px-3 border bg-zinc-800 hover:bg-zinc-900 border-zinc-800 rounded-lg">
                        Login
                    </button>

                    {/* <p className="text-xl font-bold py-1 text-zinc-800">@egrizq</p> */}
                </div>
            </nav>
            <main className="container mx-auto pt-20 text-zinc-800">
                <div className="flex justify-center">
                    <div className="flex flex-col w-6/12">

                        <Title />

                        <form onSubmit={onSubmit}
                            className="flex flex-row justify-center drop-shadow-xl pt-12">
                            <div>
                                <select id="social"
                                    onChange={(e) => setSosmed(e.target.value)}
                                    className="border hover:bg-zinc-900 border-zinc-800 bg-zinc-800 text-gray-900 text-sm rounded-l-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5  dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500">
                                    <option value="tw">X / Twitter</option>
                                    <option value="sp">Spotify</option>
                                    <option value="ig">Instagram</option>
                                    <option value="yt">Youtube</option>
                                </select>
                            </div>

                            <input type="text"
                                className="border-t border-l border-b border-zinc-800 py-1 px-2 w-7/12"
                                placeholder="Link"
                                onChange={(e) => setUrl(e.target.value)}
                            />

                            <ButtonStatus url={url} />
                        </form>

                        <div className="flex justify-center pt-7">
                            <button className="p-2 text-sm border-t border-l border-b border-zinc-300 rounded-l-md">
                                +
                            </button>
                            <div className="flex w-4/12">
                                <select onChange={(e) => setCategory(e.target.value)}
                                    className="border text-gray-800 text-sm border-zinc-300  rounded-r-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5  dark:placeholder-gray-400 dark:text-black dark:focus:ring-blue-500 dark:focus:border-blue-500">
                                    <option value="#">Pilih kategori</option>
                                    <option value="self">Self Development</option>
                                    <option value="story">Story</option>
                                </select>
                            </div>
                        </div>

                        <div className="flex justify-center pt-4">
                            {/* <ControlSosmed /> */}

                            {returnedSosmed}
                        </div>
                    </div>

                    <PopupRegister
                        isOpen={registerPopUp}
                        onClose={closeRegister}
                    />

                    <PopupLogin
                        isOpen={loginPopUp}
                        onClose={togglePopupLogin}
                        onRegister={togglePopupRegister}
                    />
                </div>
            </main>


        </>
    )
}