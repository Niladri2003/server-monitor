"use client";
import React, {useState} from "react";
import { Label } from "../components/Ui/Label.jsx";
import { Input } from "../components/Ui/Input.jsx";
import { cn } from "../utils/utils.js";
import {
    IconBrandGithub,
    IconBrandGoogle,
    IconBrandOnlyfans,
} from "@tabler/icons-react";
import Logo from "../assets/logo.jpg"
export function Login() {


    const [formData, setFormData] = useState({
        email: '',
        password: '',
    });

    // const toast = useToast();
    // const navigate = useNavigate(); // Initialize useNavigate

    const handleChange = (e) => {
        const { name, value } = e.target;
        setFormData({
            ...formData,
            [name]: value
        });
    };


    const handleSubmit = (e) => {
        e.preventDefault();
        console.log("Form submitted",formData);
    };
    return (
        (<div className={"h-screen w-screen bg-black flex items-center justify-center"}>
            <div
                className="max-w-md bg-black w-full mx-auto rounded-none border-[0.1px] border-white md:rounded-2xl p-4 md:p-8 shadow-input ">
                <div className={"w-full items-center flex justify-center mt-3 mb-2"}>
                    <img src={Logo} className={"h-16  rounded-full "}/>
                </div>
                <h2 className="font-bold text-xl text-neutral-800 dark:text-neutral-200">
                    Welcome Back to Sysmos
                </h2>
                <p className="text-neutral-600 text-sm max-w-sm mt-2 dark:text-neutral-300">
                    Sysmos is simple and easy to use server monitoring system
                </p>
                <form className="my-8" onSubmit={handleSubmit}>

                    <LabelInputContainer className="mb-4">
                        <Label htmlFor="email">Email Address</Label>
                        <Input id="email" placeholder="projectmayhem@fc.com" type="email" value={formData.email} onChange={handleChange} />
                    </LabelInputContainer>
                    <LabelInputContainer className="mb-4">
                        <Label htmlFor="password">Password</Label>
                        <Input id="password" placeholder="••••••••" type="password" value={formData.password} onChange={handleChange}/>
                    </LabelInputContainer>
                    <button
                        className="bg-gradient-to-br relative group/btn from-black dark:from-zinc-900 dark:to-zinc-900 to-neutral-600 block dark:bg-zinc-800 w-full text-white rounded-md h-10 font-medium shadow-[0px_1px_0px_0px_#ffffff40_inset,0px_-1px_0px_0px_#ffffff40_inset] dark:shadow-[0px_1px_0px_0px_var(--zinc-800)_inset,0px_-1px_0px_0px_var(--zinc-800)_inset]"
                        type="submit">
                        Login &rarr;
                        <BottomGradient/>
                    </button>

                    <div
                        className="bg-gradient-to-r from-transparent via-neutral-300 dark:via-neutral-700 to-transparent my-8 h-[1px] w-full"/>


                </form>
            </div>

        </div>)
    );
}

const BottomGradient = () => {
    return (<>
    <span
        className="group-hover/btn:opacity-100 block transition duration-500 opacity-0 absolute h-px w-full -bottom-px inset-x-0 bg-gradient-to-r from-transparent via-cyan-500 to-transparent"/>
        <span
            className="group-hover/btn:opacity-100 blur-sm block transition duration-500 opacity-0 absolute h-px w-1/2 mx-auto -bottom-px inset-x-10 bg-gradient-to-r from-transparent via-indigo-500 to-transparent"/>
    </>);
};

const LabelInputContainer = ({
                                 children,
                                 className
                             }) => {
    return (
        (<div className={cn("flex flex-col space-y-2 w-full", className)}>
            {children}
        </div>)
    );
};
