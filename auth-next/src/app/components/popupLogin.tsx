import { SubmitHandler, useForm } from "react-hook-form";
import { FormFieldslogin, response, schemaLogin } from "../auth/definitions";
import { zodResolver } from "@hookform/resolvers/zod";
import InputField from "./input";

interface checkStatus {
    isOpen: boolean
    onClose: () => void;
    onRegister: () => void;
}

const PopupLogin = ({ isOpen, onClose, onRegister }: checkStatus) => {
    if (!isOpen) return null;

    const handleOverlayClick = (e: React.MouseEvent<HTMLDivElement>) => {
        if (e.target === e.currentTarget) {
            onClose();
        }
    };

    const {
        register,
        handleSubmit,
        setError,
        formState: {
            errors,
            isSubmitting
        }
    } = useForm<FormFieldslogin>({
        resolver: zodResolver(schemaLogin),
    })

    const onSubmit: SubmitHandler<FormFieldslogin> = async (data) => {
        try {
            const response = await fetch("http://localhost:8000/login", {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(data)
            })
            if (response.ok) {
                const result = await response.json()
                console.log("success", result)

            } else {
                const errorData: response = await response.json()
                throw new Error(JSON.stringify(errorData))
            }
        } catch (error) {
            console.log(error)
            const errorMessage = (JSON.parse((error as Error).message) as response).Message;
            setError("root", {
                type: 'manual',
                message: errorMessage
            })
        }
    }

    return (
        <div className="fixed w-full inset-0 bg-gray-600 bg-opacity-50 flex justify-center items-center z-50"
            onClick={handleOverlayClick}>
            <div className="bg-white p-8 rounded-lg shadow-lg w-96">

                <h2 className="text-2xl font-bold mb-4">Login</h2>
                <form onSubmit={handleSubmit(onSubmit)}
                    className="space-y-3">
                    <InputField
                        label="Username"
                        id="username"
                        register={register}
                        error={errors.username} />
                    <InputField
                        label="Password"
                        id="password"
                        register={register}
                        error={errors.password} />

                    <div className="flex flex-col items-center justify-between pt-3 space-y-2">

                        <div className="w-full">
                            <button type="submit"
                                className="border w-full py-1 px-2 bg-zinc-800 text-white rounded-md hover:bg-zinc-900 hover:text-white">
                                {isSubmitting ? "Wait..." : "Submit"}
                            </button>
                            {errors.root &&
                                <div className="text-red-500 text-sm">{errors.root.message}</div>}
                        </div>

                        <div className="">
                            Belum punya akun?
                            <button className="px-1 font-bold" onClick={onRegister} >
                                Daftar</button>
                        </div>
                    </div>
                </form>
            </div>
        </div>
    );
};

export default PopupLogin;