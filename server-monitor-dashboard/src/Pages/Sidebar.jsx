"use client";
import { useState } from "react";
import { Sidebar, SidebarBody, SidebarLink } from "../components/sidebar/Sidebar.jsx";
import {
    IconArrowLeft,
    IconBrandTabler,
    IconSettings,
    IconUserBolt,
} from "@tabler/icons-react";
import Logoimg from "../assets/logo.webp"
import {Link, Outlet} from "react-router-dom";
import { motion } from "framer-motion";
import { cn } from "../utils/utils.js";

export function SidebarDemo() {
    const links = [
        {
            label: "Dashboard",
            to: "/dashboard",
            icon: (
                <IconBrandTabler className="text-neutral-700 dark:text-neutral-200 h-5 w-5 flex-shrink-0" />
            ),
        },
        {
            label: "Profile",
            to: "#",
            icon: (
                <IconUserBolt className="text-neutral-700 dark:text-neutral-200 h-5 w-5 flex-shrink-0" />
            ),
        },
        {
            label: "Settings",
            to: "#",
            icon: (
                <IconSettings className="text-neutral-700 dark:text-neutral-200 h-5 w-5 flex-shrink-0" />
            ),
        },
        {
            label: "Logout",
            to: "#",
            icon: (
                <IconArrowLeft className="text-neutral-700 dark:text-neutral-200 h-5 w-5 flex-shrink-0" />
            ),
        },
    ];
    const [open, setOpen] = useState(false);
    return (
        <div
            className={cn(
                " flex flex-col md:flex-row bg-gray-100 dark:bg-neutral-800 w-full h-screen flex-1  mx-auto border border-neutral-200 dark:border-neutral-700 ",
                // for your use case, use `h-screen` instead of `h-[60vh]`

            )}>
            <Sidebar open={open} setOpen={setOpen} animate={false}>
                <SidebarBody className="justify-between gap-10">
                    <div className="flex flex-col flex-1 overflow-y-auto overflow-x-hidden">
                        <Logo/>
                        <div className="mt-8 flex flex-col gap-2">
                            {links.map((link, idx) => (
                                <SidebarLink key={idx} link={link}/>
                            ))}
                        </div>
                    </div>
                    <div>
                        <SidebarLink
                            link={{
                                label: "Manu Arora",
                                to: "#",
                                icon: (
                                    <img
                                        src="https://assets.aceternity.com/manu.png"
                                        className="h-7 w-7 flex-shrink-0 rounded-full"
                                        width={50}
                                        height={50}
                                        alt="Avatar"/>
                                ),
                            }}/>
                    </div>
                </SidebarBody>
            </Sidebar>
            <div className="flex-1 overflow-y-scroll bg-white ">
                <Outlet/>
            </div>
        </div>
    );
}

export const Logo = () => {
    return (
        <Link
            to="#"
            className="font-normal flex space-x-2 items-center text-sm text-black py-1 relative z-20">
            <img src={Logoimg} className={"h-10 rounded-full"}/>
            <motion.span
                initial={{opacity: 0}}
                animate={{opacity: 1}}
                className="font-medium text-black dark:text-white whitespace-pre">
                SysMos
            </motion.span>
        </Link>
    );
};

export const LogoIcon = () => {
    return (
        <Link
            to="#"
            >
           <img src={Logoimg} />
        </Link>
    );
};
