import Head from "next/head";
import styles from "./layout.module.css";
import utilStyles from "../styles/util.module.css";

const name = "Shin Code"
export const siteTitle = "FLUXUS LOCUS BLOG";

function Layout({children}) {
    return (
        <div className={styles.container}>
            <Head>
                <link rel = "icon" href="/favicon.ico" />
            </Head>
            <header className={styles.header}>
                <img src = "/images/profile.png" className={utilStyles.borderCircle}/>
                <h1 className={utilStyles.heading2Xl}>{name}</h1>
            </header>
            <main>{children}</main>
        </div>
    );
}

export default Layout;
