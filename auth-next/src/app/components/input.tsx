import { FieldError, UseFormRegister } from "react-hook-form";

interface InputFieldProps {
    label: string;
    id: string;
    register: UseFormRegister<any>;
    error?: FieldError;
}

const InputField = ({ label, id, register, error }: InputFieldProps) => {
    return (
        <div>
            <span>{label}</span>
            <input className="w-full border border-black rounded-md py-1 px-2"
                type={label.toLowerCase().includes("password") ? "password" : "text"}
                placeholder={label}
                {...register(id)}
            />
            {error && <div className="text-red-500 text-sm">{error.message}</div>}
        </div>
    )
}

export default InputField;