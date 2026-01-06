import type {ReactNode} from 'react';
import clsx from 'clsx';
import Heading from '@theme/Heading';
import styles from './styles.module.css';
import { Zap, Cpu, Share2, ShieldCheck } from 'lucide-react';

type FeatureItem = {
  title: string;
  Icon: React.ComponentType<any>;
  iconColor: string;
  description: ReactNode;
};

const FeatureList: FeatureItem[] = [
  {
    title: 'Zero Runtime Dependencies',
    Icon: Zap,
    iconColor: '#0070f3',
    description: (
      <>
        Stop forcing Node.js or Python on your team. Freight is a single, static Go binary that works out-of-the-box on macOS, Linux, and Windows.
      </>
    ),
  },
  {
    title: 'Conductor & Railcar Architecture',
    Icon: Cpu,
    iconColor: '#0070f3',
    description: (
      <>
        The Conductor is a high-performance binary at your repo root that orchestrates events, while the Railcar (railcar.json) is your unified, declarative manifest.
      </>
    ),
  },
  {
    title: 'Effortless Team Portability',
    Icon: Share2,
    iconColor: '#0070f3',
    description: (
      <>
        Commit the Conductor binary directly to your repository. New team members and CI/CD pipelines instantly inherit your full suite of hooks without installing global tools.
      </>
    ),
  },
  {
    title: 'One-Step Engineering Safety',
    Icon: ShieldCheck,
    iconColor: '#0070f3',
    description: (
      <>
        Run <code>freight init</code> to bootstrap your repository. Freight automatically rewires every Git hook to the Conductor, ensuring your quality gates are always enforced.
      </>
    ),
  },
];

function Feature({title, Icon, iconColor, description}: FeatureItem) {
  return (
    <div className={clsx('col col--6')}>
      <div className={styles.featureCard}>
        <div className="text--center">
          <Icon className={styles.featureIcon} size={48} color={iconColor} strokeWidth={1.5} />
        </div>
        <div className="text--center padding-horiz--md">
          <Heading as="h3">{title}</Heading>
          <p>{description}</p>
        </div>
      </div>
    </div>
  );
}

export default function HomepageFeatures(): ReactNode {
  return (
    <section className={styles.features}>
      <div className="container">
        <div className="row">
          {FeatureList.map((props, idx) => (
            <Feature key={idx} {...props} />
          ))}
        </div>
      </div>
    </section>
  );
}
