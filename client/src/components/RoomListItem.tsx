import React from 'react'

type Props = {
  id: string
  name: string
  onJoin: () => void
}

const RoomListItem: React.FC<Props> = ({ id, name, onJoin }) => {
  return (
    <div className="card bg-base-100 shadow">
      <div className="card-body">
        <h3 className="card-title">{name}</h3>
        <p className="opacity-70">ID: {id}</p>
        <div className="card-actions justify-end">
          <button className="btn btn-secondary" onClick={onJoin}>Join</button>
        </div>
      </div>
    </div>
  )
}

export default RoomListItem


