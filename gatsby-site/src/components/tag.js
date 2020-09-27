import React from 'react'

const TAGS_COLORS = {
  typescript: 'is-primary',
  javascript: 'is-warning',
  webpack: 'is-link',
  网络: 'is-info',
  浏览器: 'is-success',
  'Node.js': 'is-danger',
  摘抄: 'is-primary',
  英语: 'is-warning',
  面试: 'is-info',
  todo: 'is-danger',
  算法与数据结构: 'is-success',
  阅读写作: 'is-primary',
  编辑器: 'is-link',
  工程化: 'is-danger',
  react: 'is-warning',
  设计模式: 'is-danger'
}

const Tags = ({ tags = [] }) => {
  return (
    <div className='tags'>
      {tags.map((tag, index) => (
        <span key={index} className={`tag ${TAGS_COLORS[tag]}`}>
          <strong>{tag}</strong>
        </span>
      ))}
    </div>
  )
}

export default Tags