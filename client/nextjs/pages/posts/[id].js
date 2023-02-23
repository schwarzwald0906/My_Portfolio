import Head from 'next/head'
import Layout from '../../components/layout'
import { getAllPostIds, getPostData } from '../../lib/post'
import utilStyles from '../../styles/util.module.css'

export async function getStaticPaths() {
  const paths = getAllPostIds()

  return {
    paths,
    fallback: false, //404エラー設定
  }
}

//ページに応じた記事のタイトルなどをparamsなどで引っ張ってくる。
export async function getStaticProps({ params }) {
  const postData = await getPostData(params.id)

  return {
    props: {
      postData,
    },
  }
}

export default function Post({ postData }) {
  return (
    <Layout>
      <Head>
        <title>{postData.title}</title>
      </Head>
      <article>
        <h1 className={utilStyles.headingX1}>{postData.title}</h1>
        <div className={utilStyles.lightText}>{postData.date}</div>
        <div dangerouslySetInnerHTML={{ __html: postData.blogContentHTML }} />
      </article>
    </Layout>
  )
}
