import { typeURL } from "../auth/definitions"

const ButtonStatus = ({ url }: typeURL) => {
    if (url === '') {
        return (
            <>
                <button type="submit" disabled
                    className="p-2 w-1/12 border-t border-r border-b border-zinc-800 rounded-r-md bg-zinc-800 opacity-50 text-white text-sm">
                    Add
                </button>
            </>
        )
    }

    return (
        <>
            <button type="submit"
                className="p-2 w-1/12 border-t border-r border-b border-zinc-800 rounded-r-md bg-zinc-800 hover:bg-zinc-900 text-white text-sm">
                Add
            </button>
        </>
    )
}

export default ButtonStatus