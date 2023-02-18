import Head from 'next/head'
import Link from "next/link";
import styles from "../styles/Home.module.css";
import Layout from '../components/layout';

import utilStyle from "../styles/util.module.css";

export default function Home() {
  return (
  <Layout>
    <section className={utilStyle.headingMd}>
      <p>
        私はフルスタックエンジニアです。
      </p>
    </section>

    <section className={`${utilStyle.headingMd} ${utilStyle.padding1px}`}>
      <h2>📝エンジニアのブログ</h2>
      <div className={styles.grid}>
        <article>
          <Link href="/">
            <img 
              src="/images/thumbnail01.jpg"
              className={styles.thumbnailImage}
            />
          </Link>
          <br/>
          <Link legacyBehavior href="/">
            <a className={utilStyle.boldText}>
              SSGとSSRの使い分け
            </a>
          </Link>
          <br/>
          <small className={utilStyle.lightText}>
            February 23, 2023
          </small>
        </article>
        <article>
          <Link href="/">
            <img 
              src="/images/thumbnail01.jpg"
              className={styles.thumbnailImage}
            />
          </Link>
          <br/>
          <Link legacyBehavior href="/">
            <a className={utilStyle.boldText}>
              SSGとSSRの使い分け
            </a>
          </Link>
          <br/>
          <small className={utilStyle.lightText}>
            February 23, 2023
          </small>
        </article>    
        <article>
          <Link href="/">
            <img 
              src="/images/thumbnail01.jpg"
              className={styles.thumbnailImage}
            />
          </Link>
          <br/>
          <Link legacyBehavior href="/">
            <a className={utilStyle.boldText}>
              SSGとSSRの使い分け
            </a>
          </Link>
          <br/>
          <small className={utilStyle.lightText}>
            February 23, 2023
          </small>
        </article>
        <article>
          <Link href="/">
            <img 
              src="/images/thumbnail01.jpg"
              className={styles.thumbnailImage}
            />
          </Link>
          <br/>
          <Link legacyBehavior href="/">
            <a className={utilStyle.boldText}>
              SSGとSSRの使い分け
            </a>
          </Link>
          <br/>
          <small className={utilStyle.lightText}>
            February 23, 2023
          </small>
        </article> 
      </div>
    </section>
  </Layout>
  );
}
