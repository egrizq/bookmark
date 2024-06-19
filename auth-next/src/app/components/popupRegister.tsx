import { SubmitHandler, useForm } from "react-hook-form";
import { bodyRequest, FormFieldsSignUp, response, schemaSignUp } from "../auth/definitions";
import { zodResolver } from "@hookform/resolvers/zod";
import InputField from "./input";
import { useRouter } from "next/navigation";
import { useState } from "react";

interface checkStatus {
    isOpen: boolean
    onClose: () => void;
}

const PopupRegister = ({ isOpen, onClose }: checkStatus) => {
    if (!isOpen) return null;

    const router = useRouter() // router
    const [checkResponse, setCheckResponse] = useState(false);

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
    } = useForm<FormFieldsSignUp>({
        resolver: zodResolver(schemaSignUp),
    })

    const onSubmit: SubmitHandler<FormFieldsSignUp> = async (data) => {
        try {
            const responseData: bodyRequest = {
                username: data.username,
                password: data.password,
                email: data.email
            }

            const response = await fetch("http://localhost:8000/signup", {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(responseData)
            })

            if (response.ok) {
                setCheckResponse(true)
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

    // redirect
    console.log("check", checkResponse)
    if (checkResponse) {
        router.push("/login")
    }

    return (
        <div className="fixed w-full inset-0 bg-gray-600 bg-opacity-50 flex justify-center items-center z-50"
            onClick={handleOverlayClick}>
            <div className="bg-white p-8 rounded-lg shadow-lg w-96">

                <h2 className="text-2xl font-bold mb-4">Register</h2>

                <form onSubmit={handleSubmit(onSubmit)}
                    className="space-y-3">
                    <InputField
                        label="Username"
                        id="username"
                        register={register}
                        error={errors.username} />

                    <InputField
                        label="Email"
                        id="email"
                        register={register}
                        error={errors.email} />

                    <InputField
                        label="Password"
                        id="password"
                        register={register}
                        error={errors.password} />

                    <InputField
                        label="Password Confirmation"
                        id="passwordConf"
                        register={register}
                        error={errors.passwordConf} />

                    <div className="flex flex-col items-center justify-between pt-3 space-y-2">
                        <div className="w-full">
                            <button type="submit"
                                className="border w-full py-1 px-2 bg-zinc-800 text-white rounded-md hover:bg-zinc-900 hover:text-white">
                                {isSubmitting ? "Wait..." : "Submit"}
                            </button>
                            {errors.root &&
                                <div className="text-red-500 text-sm">{errors.root.message}</div>}
                        </div>
                    </div>

                </form>
            </div>
        </div>
    );
};

export default PopupRegister;