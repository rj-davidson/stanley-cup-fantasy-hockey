import React from 'react';
import TeamRoster, { Props as TeamRosterProps } from '@/components/team/team';

interface Props extends TeamRosterProps {
  index: number;
}

const SortableTeamRoster: React.FC<Props> = ({ index, ...rest }) => {
  const { owner, roster } = rest;
  const totalPoints = roster.reduce((sum, player) => sum + player.Points, 0);

  return (
    <div
      style={{ marginBottom: 16 }}
      data-index={index}
      data-total-points={totalPoints}
    >
      <TeamRoster {...rest} />
    </div>
  );
};

type TeamsListProps = {
  teams: Props[];
};

const TeamsList: React.FC<TeamsListProps> = ({ teams }) => {
  const sortedTeams = [...teams].sort((a, b) => {
    const aTotalPoints = parseInt(
      a.roster.reduce((sum, player) => sum + player.Points, 0).toString(),
    );
    const bTotalPoints = parseInt(
      b.roster.reduce((sum, player) => sum + player.Points, 0).toString(),
    );
    return bTotalPoints - aTotalPoints;
  });

  return (
    <>
      {sortedTeams.map((team, key) => (
        <SortableTeamRoster key={key} {...team} />
      ))}
    </>
  );
};

export default TeamsList;
